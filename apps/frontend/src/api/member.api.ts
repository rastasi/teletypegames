import type { Member } from '../lib/interfaces/member.interface'

const index = async (): Promise<Member[]> => {
  const res = await fetch('/api/members')
  if (!res.ok) throw new Error(`HTTP ${res.status}`)
  return res.json()
}

export default { index }
