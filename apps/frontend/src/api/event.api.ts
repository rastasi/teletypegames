import type { Event } from '../lib/interfaces/event.interface'

const index = async (): Promise<Event[]> => {
  const res = await fetch('/api/events')
  if (!res.ok) throw new Error(`HTTP ${res.status}`)
  return res.json()
}

export default { index }
