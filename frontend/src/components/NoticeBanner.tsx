import type { Notice } from '../types/inventory'

type NoticeBannerProps = {
  notice: Notice | null
}

export function NoticeBanner({ notice }: NoticeBannerProps) {
  if (!notice) return null

  return (
    <div className={`notice ${notice.type}`} role="status">
      {notice.message}
    </div>
  )
}
