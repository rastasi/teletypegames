export const DATETIME_FORMAT = 'YYYY-MM-DD HH:mm'
export const DATE_FORMAT = 'YYYY-MM-DD'

export function formatDateTime(dateStr: string | Date): string {
  const d = new Date(dateStr as string)
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  const hours = String(d.getHours()).padStart(2, '0')
  const minutes = String(d.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}`
}

export function formatDate(dateStr: string | Date): string {
  const d = new Date(dateStr as string)
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}
