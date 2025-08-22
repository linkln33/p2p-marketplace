export type Listing = {
  id: number;
  seller: string;
  title: string;
  description: string;
  price: number;
  denom: string;
  status: string;
  buyer: string;
  created_at: number;
  expires_at: number;
};

export type ListListingResponse = {
  listing: Listing[];
  pagination?: {
    next_key?: string;
    total?: string;
  };
};

const LCD = process.env.NEXT_PUBLIC_LCD_URL || "http://127.0.0.1:1317";

export async function fetchListings(): Promise<Listing[]> {
  const url = `${LCD}/market/market/v1/listing`;
  const res = await fetch(url, { cache: "no-store" });
  if (!res.ok) {
    throw new Error(`Failed to fetch listings: ${res.status} ${res.statusText}`);
  }
  const data = (await res.json()) as ListListingResponse;
  return data.listing || [];
}
