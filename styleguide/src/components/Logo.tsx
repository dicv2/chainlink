import React from 'react'
import { Image } from './Image'

interface IProps {
  src: string
  width?: number
  height?: number
  alt?: string
}

export const Logo = (props: IProps) => {
  return <Image {...props} />
}
