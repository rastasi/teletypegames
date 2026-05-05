export interface WikiPage {
  id: number
  path: string
  title: string
  description: string
  updatedAt: string
  createdAt: string
  locale: string
}

export interface WikiPageWithContent extends WikiPage {
  content: string
}

export interface WikiPageContent extends WikiPage {
  render: string
}
