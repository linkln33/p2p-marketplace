import { fetchListings } from "@/lib/api";

export default async function Home() {
  let listings: Awaited<ReturnType<typeof fetchListings>> = [];
  let error: string | null = null;
  try {
    listings = await fetchListings();
  } catch (e: any) {
    error = e?.message || "Failed to load listings";
  }

  return (
    <div className="min-h-screen p-8">
      <main className="max-w-5xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">P2P Marketplace</h1>
        {error ? (
          <div className="rounded border border-red-300 bg-red-50 p-4 text-red-700">
            <p className="font-semibold">Error loading listings</p>
            <p className="text-sm mt-1">{error}</p>
            <p className="text-xs mt-2 text-red-600/80">
              Ensure the node REST API is running at http://127.0.0.1:1317 and the market module is enabled.
            </p>
          </div>
        ) : listings.length === 0 ? (
          <div className="rounded border p-6 bg-gray-50">
            <p className="text-gray-700">No listings yet.</p>
            <p className="text-gray-500 text-sm">Create one from the UI once wallet actions are wired.</p>
          </div>
        ) : (
          <ul className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {listings.map((l) => (
              <li key={l.id} className="rounded-lg border p-4">
                <div className="flex items-start justify-between">
                  <h2 className="text-lg font-semibold mr-3 line-clamp-1">{l.title}</h2>
                  <span className="text-xs rounded px-2 py-0.5 border bg-gray-50 text-gray-600">{l.status}</span>
                </div>
                <p className="text-sm text-gray-600 mt-1 line-clamp-2">{l.description}</p>
                <div className="flex items-center justify-between mt-3 text-sm">
                  <span className="font-medium">{l.price} {l.denom}</span>
                  <span className="text-gray-500 truncate max-w-[50%]">Seller: {l.seller}</span>
                </div>
              </li>
            ))}
          </ul>
        )}
      </main>
    </div>
  );
}
