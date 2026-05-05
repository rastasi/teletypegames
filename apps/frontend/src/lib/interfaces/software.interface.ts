export interface Release {
  version: string
  htmlFolderPath?: string
  cartridgePath?: string
  sourcePath?: string
  docsFolderPath?: string
  UpdatedAt: string
}

export interface Software {
  name: string
  title: string
  desc: string
  status: string
  author: string
  platform: string
  license?: string
  story?: string
  externalLinks?: { label: string; url: string }[]
}

export interface SoftwareEntry {
  software: Software
  releases: Release[]
}
