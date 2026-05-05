import type { WikiPage, WikiPageWithContent, WikiPageContent } from '../lib/interfaces/wiki.interface'

const BASE = import.meta.env.VITE_WIKI_BASE || 'https://wiki.teletype.hu'
const TOKEN = import.meta.env.WEBAPP_WIKIJS_TOKEN

function buildHeaders(): Record<string, string> {
  const h: Record<string, string> = { 'Content-Type': 'application/json', 'Accept': 'application/json' }
  if (TOKEN) h['Authorization'] = `Bearer ${TOKEN}`
  return h
}

async function gql(query: string): Promise<any> {
  const res = await fetch(`${BASE}/graphql`, {
    method: 'POST',
    headers: buildHeaders(),
    body: JSON.stringify({ query }),
  })
  const json = await res.json()
  if (json.errors) throw new Error(json.errors.map((e: any) => e.message).join(', '))
  return json.data
}

const listBlogPages = async (): Promise<WikiPageWithContent[]> => {
  const data = await gql(`{
    pages {
      list(orderBy: CREATED, orderByDirection: DESC, tags: ["blog"]) {
        id path title description updatedAt createdAt locale
      }
    }
  }`)
  const pages: any[] = data?.pages?.list ?? []

  const pagesWithContent = await Promise.all(pages.map(async (p: any) => {
    try {
      const contentData = await gql(`{ pages { single(id: ${p.id}) { content } } }`)
      return { ...p, content: contentData?.pages?.single?.content ?? '' }
    } catch {
      return { ...p, content: '' }
    }
  }))

  return pagesWithContent.map((p: any): WikiPageWithContent => ({
    id: p.id,
    path: p.path,
    title: p.title || p.path,
    description: p.description ?? '',
    content: p.content ?? '',
    updatedAt: p.updatedAt,
    createdAt: p.createdAt,
    locale: p.locale,
  }))
}

const getBlogPage = async (slug: string): Promise<WikiPageContent | null> => {
  const data = await gql(`{
    pages {
      list(orderBy: CREATED, orderByDirection: DESC, tags: ["blog"]) {
        id path
      }
    }
  }`)
  const pages: any[] = data?.pages?.list ?? []

  const matched = pages.find((p: any) => {
    const pageSlug = p.path.startsWith('blog/') ? p.path.replace('blog/', '') : p.path
    return pageSlug === slug
  })

  if (!matched) return null

  const singleData = await gql(`{
    pages {
      single(id: ${matched.id}) {
        id path title description render updatedAt createdAt locale
      }
    }
  }`)
  return singleData?.pages?.single ?? null
}

const listHowtoPages = async (): Promise<WikiPage[]> => {
  const data = await gql(`{
    pages {
      list(orderBy: UPDATED, orderByDirection: DESC, tags: ["howto"]) {
        id path title description updatedAt createdAt locale
      }
    }
  }`)
  const pages: any[] = data?.pages?.list ?? []
  return pages.map((p: any): WikiPage => ({
    id: p.id,
    path: p.path,
    title: p.title || p.path,
    description: p.description ?? '',
    updatedAt: p.updatedAt,
    createdAt: p.createdAt,
    locale: p.locale,
  })).slice(0, 30)
}

export { BASE as WIKI_BASE }
export default { listBlogPages, getBlogPage, listHowtoPages }
