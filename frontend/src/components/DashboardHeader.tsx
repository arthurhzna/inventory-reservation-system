type DashboardHeaderProps = {
  apiBaseUrl: string
}

export function DashboardHeader({ apiBaseUrl }: DashboardHeaderProps) {
  return (
    <header className="topbar">
      <div>
        <p className="eyebrow">Inventory Reservation System</p>
        <h1>Mini Dashboard</h1>
      </div>
      <div className="api-chip">
        <span></span>
        {apiBaseUrl}
      </div>
    </header>
  )
}
