package policy

import (
	"testing"
)

// BlockchainList Tests
// First, some tests where a subset (or equal) is found
func Test_Blockchain_subset_found(t *testing.T) {
	var prod_bl *BlockchainList
	var con_bl *BlockchainList

	prod1 := `[]`
	con1 := `[{"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err != nil {
				t.Errorf("Error: %v is a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	prod1 = `[{"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err != nil {
				t.Errorf("Error: %v is a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	con1 = `[{"type":"hyperledger","details":{"boot":"http://bhnetwork.com/boot"}},
             {"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err != nil {
				t.Errorf("Error: %v is a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	con1 = `[{"type":"ethereum","details":{"genesis":["http://mycompany.com/genesis"],"networkid":["http://mycompany.com/networkid"],"bootnodes":["http://mycompany.com/bootnodes"],"directory":["http://mycompany.com/directory"]}},
             {"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	prod1 = `[{"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}},
              {"type":"ethereum","details":{"genesis":["http://mycompany.com/genesis"],"networkid":["http://mycompany.com/networkid"],"bootnodes":["http://mycompany.com/bootnodes"],"directory":["http://mycompany.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err != nil {
				t.Errorf("Error: %v is a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}
}

// Second, some tests where a subset (or equal) is NOT found
func Test_Blockchain_subset_not_found(t *testing.T) {
	var prod_bl *BlockchainList
	var con_bl *BlockchainList

	prod1 := `[{"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	con1 := `[]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err == nil {
				t.Errorf("Error: %v is not a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	con1 = `[{"type":"hyperledger","details":{"boot":"http://bhnetwork.com/boot"}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err == nil {
				t.Errorf("Error: %v is not a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	con1 = `[{"type":"ethereum","details":{"genesis":["http://mycompany.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err == nil {
				t.Errorf("Error: %v is not a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	con1 = `[{"type":"ethereum","details":{"genesis":["http://mycompany.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://bhnetwork.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}},
             {"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"networkid":["http://bhnetwork.com/networkid"],"bootnodes":["http://mycompany.com/bootnodes"],"directory":["http://bhnetwork.com/directory"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err == nil {
				t.Errorf("Error: %v is not a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}

	prod1 = `[{"type":"ethereum","details":{"genesis":["http://bhnetwork.com/genesis"],"bootnodes":["http://bhnetwork.com/bootnodes"]}}]`
	if prod_bl = create_BlockchainList(prod1, t); prod_bl != nil {
		if con_bl = create_BlockchainList(con1, t); con_bl != nil {
			if err := prod_bl.Is_Subset_Of(con_bl); err == nil {
				t.Errorf("Error: %v is not a subset of %v, subset error was %v\n", prod1, con1, err)
			}
		}
	}
}
