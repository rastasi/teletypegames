import type { GiteaRepo, Commit } from '../lib/interfaces/git.interface'

const BASE = import.meta.env.VITE_GIT_BASE || 'https://git.teletypegames.org/api/v1'
const TOKEN = import.meta.env.WEBAPP_GITEA_TOKEN

function buildHeaders() {
  return { 'Authorization': `token ${TOKEN}`, 'Accept': 'application/json' }
}

const repos = async (): Promise<GiteaRepo[]> => {
  if (!TOKEN) throw new Error('Gitea API token (WEBAPP_GITEA_TOKEN) is not configured.')
  const res = await fetch(`${BASE}/repos/search?q=&private=false&limit=50`, { headers: buildHeaders() })
  if (!res.ok) throw new Error(`Gitea API responded with status ${res.status}`)
  const data = await res.json()
  return data.data || []
}

const commits = async (owner: string, name: string, htmlUrl: string): Promise<Commit[]> => {
  try {
    const res = await fetch(`${BASE}/repos/${owner}/${name}/commits?limit=10&page=1`, { headers: buildHeaders() })
    if (!res.ok) return []
    const data = await res.json()
    if (!Array.isArray(data)) return []
    return data.map((c: any): Commit => ({
      sha: c.sha,
      message: c.commit?.message ?? '',
      author: { name: c.commit?.author?.name ?? 'Unknown', email: c.commit?.author?.email ?? '' },
      date: c.commit?.author?.date ?? c.created,
      url: c.html_url,
      repo: { owner, name, html_url: htmlUrl },
    }))
  } catch {
    return []
  }
}

export default { repos, commits }
