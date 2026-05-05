import type { SoftwareEntry } from '../lib/interfaces/software.interface'

const index = async (): Promise<SoftwareEntry[]> => {
  const res = await fetch('/api/software')
  if (!res.ok) throw new Error(`HTTP ${res.status}`)
  const json = await res.json()
  return json.softwares
}

const highlighted = async (): Promise<SoftwareEntry | null> => {
  const res = await fetch('/api/software/highlighted')
  if (!res.ok) return null
  return res.json()
}

export default { index, highlighted }
