import { Server } from 'http'
import WebSocket from 'ws'
import { start as startServer, DEFAULT_TEST_PORT } from '../support/server'
import { Connection } from 'typeorm'
import { closeDbConnection, getDb } from '../database'
import createFixture from './fixtures/JobRun.fixture.json'
import updateFixture from './fixtures/JobRunUpdate.fixture.json'
import { JobRun } from '../entity/JobRun'
import { TaskRun } from '../entity/TaskRun'
import { Client, createClient, deleteClient } from '../entity/Client'

const ENDPOINT = `ws://localhost:${DEFAULT_TEST_PORT}`

const newExploreclient = (
  url: string,
  accessKey: string,
  secret: string
): Promise<WebSocket> => {
  const ws = new WebSocket(ENDPOINT, {
    headers: {
      'X-Explore-Chainlink-AccessKey': accessKey,
      'X-Explore-Chainlink-Secret': secret
    }
  })

  return new Promise((resolve: (arg0: WebSocket) => void, reject) => {
    ws.on('error', (error: Error) => {
      reject(error)
    })

    ws.on('open', () => resolve(ws))
  })
}

describe('realtime', () => {
  let server: Server
  let db: Connection
  let client: Client
  let secret: string

  beforeAll(async () => {
    server = await startServer()
    db = await getDb()
  })

  beforeEach(async () => {
    ;[client, secret] = await createClient(db, 'explore client')
  })

  afterAll(async () => {
    return Promise.all([server.close(), closeDbConnection()])
  })

  it('create a job run for valid JSON', async () => {
    expect.assertions(3)

    const ws = await newExploreclient(ENDPOINT, client.accessKey, secret)

    ws.send(JSON.stringify(createFixture))

    await new Promise(resolve => {
      ws.on('message', (data: WebSocket.Data) => {
        const result = JSON.parse(data as string)
        expect(result.status).toEqual(201)
        ws.close()
        resolve()
      })
    })

    const jobRunCount = await db.manager.count(JobRun)
    expect(jobRunCount).toEqual(1)

    const taskRunCount = await db.manager.count(TaskRun)
    expect(taskRunCount).toEqual(1)
  })

  it('can create and update a job run and task runs', async (done: any) => {
    expect.assertions(6)

    const db = await getDb()
    const assertionCallback = async () => {
      const jobRunCount = await db.manager.count(JobRun)
      expect(jobRunCount).toEqual(1)

      const taskRunCount = await db.manager.count(TaskRun)
      expect(taskRunCount).toEqual(1)

      const jr = await db.manager.findOne(JobRun, { relations: ['taskRuns'] })
      expect(jr.status).toEqual('completed')

      const tr = jr.taskRuns[0]
      expect(tr.status).toEqual('completed')
      done()
    }

    const ws = new WebSocket(ENDPOINT)
    let responses = 0
    ws.on('message', (data: any) => {
      responses += 1
      const result = JSON.parse(data)
      expect(result.status).toEqual(201)
      if (responses === 2) {
        ws.close()
        return assertionCallback()
      }
    })

    // send payloads
    ws.on('open', () => {
      ws.send(JSON.stringify(createFixture))
      ws.send(JSON.stringify(updateFixture))
    })
  })

  it('rejects malformed json events with code 422', async (done: any) => {
    expect.assertions(2)

    const ws = await newExploreclient(ENDPOINT, client.accessKey, secret)

    ws.send('{invalid json}')

    ws.on('message', async (data: any) => {
      const result = JSON.parse(data)
      expect(result.status).toEqual(422)

      const count = await db.manager.count(JobRun)
      expect(count).toEqual(0)

      ws.close()
      done()
    })
  })

  it('reject invalid authentication', async (done: any) => {
    expect.assertions(1)

    newExploreclient(ENDPOINT, client.accessKey, 'lol-no').catch(error => {
      expect(error).toBeDefined()
      done()
    })
  })
})
