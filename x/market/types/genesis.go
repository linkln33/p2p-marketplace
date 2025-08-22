package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
    return &GenesisState{
        Params:      DefaultParams(),
        ListingList: []Listing{},
    }
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
    listingIdMap := make(map[uint64]bool)
    listingCount := gs.GetListingCount()
    for _, elem := range gs.ListingList {
        if _, ok := listingIdMap[elem.Id]; ok {
            return fmt.Errorf("duplicated id for listing")
        }
        if elem.Id >= listingCount {
            return fmt.Errorf("listing id should be lower or equal than the last id")
        }
        listingIdMap[elem.Id] = true
    }

    return gs.Params.Validate()
}
