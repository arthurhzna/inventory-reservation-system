import type { InventoryItem } from '../types/inventory'

type InventoryTrackerProps = {
  inventory: InventoryItem[]
  selectedItem: InventoryItem | null
  selectedItemId: string
  isLoading: boolean
  lastUpdatedAt: Date | null
  onRefresh: () => void
  onSelectItem: (itemId: string) => void
}

export function InventoryTracker({
  inventory,
  selectedItem,
  selectedItemId,
  isLoading,
  lastUpdatedAt,
  onRefresh,
  onSelectItem,
}: InventoryTrackerProps) {
  return (
    <div className="panel inventory-panel">
      <div className="panel-header">
        <div>
          <p className="eyebrow">Live Inventory</p>
          <h2>{selectedItem?.item_name ?? 'Select item'}</h2>
        </div>
        <button
          className="secondary-button"
          type="button"
          onClick={onRefresh}
          disabled={isLoading}
        >
          {isLoading ? 'Refreshing' : 'Refresh'}
        </button>
      </div>

      <label className="field">
        Item
        <select
          value={selectedItemId}
          onChange={(event) => onSelectItem(event.target.value)}
        >
          {inventory.length === 0 ? (
            <option value={selectedItemId}>{selectedItemId}</option>
          ) : (
            inventory.map((item) => (
              <option key={item.item_id} value={item.item_id}>
                {item.item_name} ({item.item_id})
              </option>
            ))
          )}
        </select>
      </label>

      <div className="stock-grid">
        <div className="stock-metric">
          <span>Total</span>
          <strong>{selectedItem?.total_stock ?? '-'}</strong>
        </div>
        <div className="stock-metric">
          <span>Reserved</span>
          <strong>{selectedItem?.reserved_stock ?? '-'}</strong>
        </div>
        <div className="stock-metric available">
          <span>Available</span>
          <strong>{selectedItem?.available_stock ?? '-'}</strong>
        </div>
      </div>

      <div className="inventory-table">
        <div className="table-row table-heading">
          <span>Item</span>
          <span>Available</span>
        </div>
        {inventory.map((item) => (
          <button
            className="table-row table-button"
            type="button"
            key={item.item_id}
            onClick={() => onSelectItem(item.item_id)}
          >
            <span>{item.item_name}</span>
            <strong>{item.available_stock}</strong>
          </button>
        ))}
      </div>

      <p className="muted">
        Last updated: {lastUpdatedAt?.toLocaleTimeString() ?? 'Waiting'}
      </p>
    </div>
  )
}
