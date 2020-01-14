import { createSelector } from 'reselect'

const NODE_NAMES = [
  {
    address: '0x049bd8c3adc3fe7d3fc2a44541d955a537c2a484',
    name: 'Fiews',
  },
  {
    address: '0x240bae5a27233fd3ac5440b5a598467725f7d1cd',
    name: 'LinkPool',
  },
  {
    address: '0x4565300c576431e5228e8aa32642d5739cf9247d',
    name: 'Certus One',
  },
  {
    address: '0x58c69aff4df980357034ea98aad35bbf78cbd849',
    name: 'Wetez',
  },
  {
    address: '0x79c6e11be1c1ed4d91fbe05d458195a2677f14a5',
    name: 'Validation Capital',
  },
  {
    address: '0x7a9d706b2a3b54f7cf3b5f2fcf94c5e2b3d7b24b',
    name: 'LinkForest',
  },
  {
    address: '0x7e94a8a23687d8c7058ba5625db2ce358bcbd244',
    name: 'SNZPool',
  },
  {
    address: '0x89f70fa9f439dbd0a1bc22a09befc56ada04d9b4',
    name: 'Chainlink',
  },
  {
    address: '0x8c85a06eb3854df0d502b2b00169dbfb8b603bf3',
    name: 'Kaiko',
  },
  {
    address: '0x8cfb1d4269f0daa003cdea567ac8f76c0647764a',
    name: 'Simply VC',
  },
  {
    address: '0xb92ec7d213a28e21b426d79ede3c9bbcf6917c09',
    name: 'stake.fish',
  },
  {
    address: '0xf3b450002c7bc300ea03c9463d8e8ba7f821b7c6',
    name: 'Newroad',
  },
  {
    address: '0xf5a3d443fccd7ee567000e43b23b0e98d96445ce',
    name: 'Chainlayer',
  },
  {
    address: '0x992Ef8145ab8B3DbFC75523281DaD6A0981891bb',
    name: 'Figment Networks',
  },
  {
    address: '0x83dA1beEb89Ffaf56d0B7C50aFB0A66Fb4DF8cB1',
    name: 'Omniscience',
  },
  {
    address: '0x0Ce0224ba488ffC0F46bE32b333a874Eb775c613',
    name: 'Cosmostation',
  },
  {
    address: '0x64FE692be4b42F4Ac9d4617aB824E088350C11C2',
    name: 'Ztake.org',
  },
  {
    address: '0x260A96cEC05328f678754D1ACF143C8ac1DF079A',
    name: 'HashQuark',
  },
  {
    address: '0x38b6ab6B9294CCe1Ccb59c3e7D390690B4c18B1A',
    name: 'Prophet',
  },
  {
    address: '0x2Ed7E9fCd3c0568dC6167F0b8aEe06A02CD9ebd8',
    name: 'Secure Data Links',
  },
  {
    address: '0x78E76126719715Eddf107cD70f3A31dddF31f85A',
    name: 'Honeycomb.market',
  },
  {
    address: '0x24A718307Ce9B2420962fd5043fb876e17430934',
    name: 'Infinity Stones',
  },
  {
    address: '0x72f3dFf4CD17816604dd2df6C2741e739484CA62',
    name: 'Alpha Vantage',
  },
  {
    address: '0x29e3b3c76e7ae0d681bf1a6BceE1c0E7d17DBAA9',
    name: 'P2P.org',
  },
  {
    address: '0x9308B0Bd23794063423f484Cd21c59eD38898108',
    name: 'Anyblock',
  },

  // ROPSTEN

  {
    address: '0x83F00b902cbf06E316C95F51cbEeD9D2572a349a',
    name: 'LinkPool',
  },
  {
    address: '0xc2e33121d00064841e844EB068221803e140f496',
    name: 'stake.fish',
  },
  {
    address: '0x6E6F16B7C0A00A2aC1136b3aE3E4641f1FAF8D7f',
    name: 'Chainlayer',
  },
  {
    address: '0x90eeb07A0DdB176D4c60deC3a146e2289DCB2674',
    name: 'Cosmostation',
  },
  {
    address: '0x1041F70920ec1d0A10C53a3cA1235d5971d0a0ED',
    name: 'Prophet',
  },
  {
    address: '0x0D31C381c84d94292C07ec03D6FeE0c1bD6e15c1',
    name: 'Simply VC',
  },
  {
    address: '0x4a3FBbB385b5eFEB4BC84a25AaADcD644Bd09721',
    name: 'Honeycomb.market',
  },
  {
    address: '0x3a44681bdEDa78F2C901161d9ADCcFd23b6DaA36',
    name: 'Newroad',
  },
  {
    address: '0x1948C20CC492539968BB9b041F96D6556B4b7001',
    name: 'Fiews',
  },
  {
    address: '0x4105d850E9Aea215f9350C9E46Bb73FC0448C20a',
    name: 'LinkForest',
  },
  {
    address: '0xfe08369B2021c194C86cb05aF0D15f837561E09b',
    name: 'Figment Networks',
  },
  {
    address: '0xfc1dA11c2477a6c398E9e595ACE8a09064636e9D',
    name: 'Validation Capital',
  },
  {
    address: '0x80DAD789487f9EF2d139C8AFD18a6Ac5f5530e28',
    name: 'Wetez',
  },
  {
    address: '0xA3Ce768F041d136E8d57fD24372E5fB510b797ec',
    name: 'Certus One',
  },
  {
    address: '0x83dA1beEb89Ffaf56d0B7C50aFB0A66Fb4DF8cB1',
    name: 'Omniscience',
  },
  {
    address: '0xAB59b29D018a522da92761af759BEa2c272C0c41',
    name: 'Ztake.org',
  },
  {
    address: '0xa0BfFBdf2c440D6c76af13c30d9B320F9d2DeA6A',
    name: 'Secure Data Links',
  },
  {
    address: '0xc99B3D447826532722E41bc36e644ba3479E4365',
    name: 'Chainlink',
  },
  {
    address: '0x948374A346fDf2F73545864fAb6e3488Ed125961',
    name: 'Infinity Stones',
  },
  {
    address: '0x2c3E0524d8B9601916A29A114a150610EB1653ce',
    name: 'SNZPool',
  },
  {
    address: '0xb36d3709e22f7c708348e225b20b13ea546e6d9c',
    name: 'Alpha Vantage',
  },
  {
    address: '0xa8EE2a4a0F010E66A0f1C4321865D3dbeE3070Cb',
    name: 'Kaiko',
  },
  {
    address: '0x85aEace84a130bC1AcCcE2a9F4F933F6765b0B9B',
    name: 'Everstake',
  },
]

const oracles = state => state.aggregation.oracles
const oracleResponse = state => state.aggregation.oracleResponse
const currentAnswer = state => state.aggregation.currentAnswer
const contractAddress = state => state.aggregation.contractAddress
const pendingAnswerId = state => state.aggregation.pendingAnswerId

const oraclesList = createSelector([oracles], list => {
  if (!list) return []

  const names = {}

  NODE_NAMES.forEach(n => {
    names[n.address.toUpperCase()] = n.name
  })

  const result = list.map(a => {
    return {
      address: a,
      name: names[a.toUpperCase()] || 'Unknown',
      type: 'oracle',
    }
  })

  return result
})

const networkGraphNodes = createSelector(
  [oraclesList, contractAddress],
  (list, address) => {
    if (!list) return []

    let result = [
      {
        type: 'contract',
        name: 'Aggregation Contract',
        address,
      },
      ...list,
    ]

    result = result.map((a, i) => {
      return { ...a, id: i }
    })

    return result
  },
)

const networkGraphState = createSelector(
  [oracleResponse, currentAnswer],
  (list, answer) => {
    if (!list) return []

    const contractData = {
      currentAnswer: answer,
      type: 'contract',
    }

    return [...list, contractData]
  },
)

const oraclesData = createSelector(
  [oraclesList, oracleResponse, pendingAnswerId],
  (list, response, pendingAnswerId) => {
    if (!list) return []

    const data = list.map((o, id) => {
      const state = response && response.filter(r => r.sender === o.address)[0]
      const isFulfilled = state && state.answerId >= pendingAnswerId
      return { ...o, ...state, id, isFulfilled }
    })

    return data
  },
)

export { oraclesList, networkGraphNodes, networkGraphState, oraclesData }
