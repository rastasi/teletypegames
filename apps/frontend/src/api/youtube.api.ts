import type { YoutubeVideo } from '../lib/interfaces/youtube.interface'

const API_KEY = import.meta.env.YOUTUBE_API_KEY
const CHANNEL_ID = import.meta.env.YOUTUBE_CHANNEL_ID

const latestVideo = async (): Promise<YoutubeVideo | null> => {
  if (!API_KEY || !CHANNEL_ID) return null

  const searchRes = await fetch(
    `https://www.googleapis.com/youtube/v3/search?key=${API_KEY}&channelId=${CHANNEL_ID}&part=snippet&order=date&maxResults=1&type=video`
  )
  if (!searchRes.ok) throw new Error(`API error: ${searchRes.status}`)
  const searchData = await searchRes.json()

  if (!searchData.items?.length) throw new Error('No videos found on this channel.')

  const item = searchData.items[0]
  const videoId = item.id.videoId

  const statsRes = await fetch(
    `https://www.googleapis.com/youtube/v3/videos?key=${API_KEY}&id=${videoId}&part=statistics`
  )
  const statsData = await statsRes.json()
  const views = statsData.items?.[0]?.statistics?.viewCount ?? null

  return {
    id: videoId,
    title: item.snippet.title,
    publishDate: item.snippet.publishedAt,
    viewCount: views ? views.toString() : '',
  }
}

export default { latestVideo }
