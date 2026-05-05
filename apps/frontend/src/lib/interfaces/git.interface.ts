export interface GiteaRepo {
  owner: { login: string }
  name: string
  description: string
  language?: string
  html_url: string
  updated_at: string
}

export interface Commit {
  sha: string
  message: string
  author: { name: string; email: string }
  date: string
  url: string
  repo: { owner: string; name: string; html_url: string }
}
