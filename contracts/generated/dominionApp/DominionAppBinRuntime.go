package dominionApp

// DominionAppBinRuntime is the runtime part of the compiled bytecode used for deploying new contracts.
var DominionAppBinRuntime = "608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610c2b565b610045565b005b6000604051806020016040528061009f8680606001906100659190610d73565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506101ac92505050565b9052905060006100ae8261025e565b9050600060405180602001604052806100d08880606001906100659190610d73565b9052905060006100df8261025e565b9050600060405180602001604052806101018980606001906100659190610d73565b9052905060006101108261025e565b80515190915060ff1687146101405760405162461bcd60e51b815260040161013790610cee565b60405180910390fd5b600261014f60408c018c610d25565b90501461016e5760405162461bcd60e51b815260040161013790610cbe565b845160600151815160200151600881111561018557fe5b8151811061018f57fe5b60200260200101516101a057600080fd5b50505050505050505050565b60606000825167ffffffffffffffff811180156101c857600080fd5b506040519080825280602002602001820160405280156101f2578160200160208202803683370190505b50905060005b83518163ffffffff16101561025757838163ffffffff168151811061021957fe5b602001015160f81c60f81b828263ffffffff168151811061023657fe5b6001600160f81b0319909216602092830291909101909101526001016101f8565b5092915050565b610266610b07565b600061027183610309565b9050600061027e84610334565b6040805160028082526060820190925291925060009190816020015b6102a2610b40565b81526020019060019003908161029a57905050905060005b60028110156102d5576102cc86610353565b506001016102ba565b5060006102e186610372565b6040805160808101825295865260208601949094529284019190915250606082015292915050565b610311610b87565b600061032d6103288461032386610391565b6103ce565b610404565b9392505050565b61033c610bb0565b600061032d61034e8461032386610391565b61051b565b61035b610b40565b600061032d61036d8461032386610391565b6105a2565b61037a610bc3565b600061032d61038c8461032386610391565b61079d565b60008082600001516000815181106103a557fe5b602002602001015190506103c48360000151600185600001515161081c565b9092525060f81c90565b606060006103e5846000015160008560ff1661081c565b845180519192506103fa9160ff86169061081c565b8452905092915050565b61040c610b87565b6040805160088082526101208201909252600091602082016101008036833701905050905060005b60038451038160ff16101561048c57610465848260030160ff168151811061045857fe5b60200260200101516108c7565b828260ff168151811061047457fe5b91151560209283029190910190910152600101610434565b506040518060800160405280846000815181106104a557fe5b602002602001015160f81c60ff168152602001846001815181106104c557fe5b602002602001015160f81c60ff1660088111156104de57fe5b60088111156104e957fe5b815260200161050a6104fe866002600361081c565b60008151811061045857fe5b151581526020019190915292915050565b610523610bb0565b6040805160066020820181815261010083018452600093839290830160c0803683375050509052905060005b600660ff8216101561025757838160ff168151811061056a57fe5b602002602001015160f81c82600001518260ff168151811061058857fe5b60ff9092166020928302919091019091015260010161054f565b6105aa610b40565b6000826000815181106105b957fe5b602002602001015160f81c905060006105e26105dd8560018560010160ff1661081c565b6108d0565b90506105f6848360010160ff16865161081c565b935060008460008151811061060757fe5b602002602001015160f81c9050600061062b6105dd8760018560010160ff1661081c565b905061063f868360010160ff16885161081c565b955060008660008151811061065057fe5b602002602001015160f81c905060006106746105dd8960018560010160ff1661081c565b9050610688888360010160ff168a5161081c565b975060008860008151811061069957fe5b602002602001015160f81c905060006106bd6105dd8b60018560010160ff1661081c565b90506106d18a8560010160ff168c5161081c565b60408051600480825260a08201909252919b50600091906020820160808036833701905050905060008b60008151811061070757fe5b602002602001015160f81c905060005b8160ff168160ff16101561076b578c8160010160ff168151811061073757fe5b602002602001015160f81c838260ff168151811061075157fe5b60ff90921660209283029190910190910152600101610717565b50506040805160a081018252988952602089019690965294870192909252506060850152506080830152509392505050565b6107a5610bc3565b815160608080601460ff8516106107c6576107c3866000601461081c565b92505b602860ff8516106107e1576107de866014602861081c565b91505b603c60ff8516106107fc576107f9866028603c61081c565b90505b604080516060810182529384526020840192909252908201529392505050565b6060600083830367ffffffffffffffff8111801561083957600080fd5b50604051908082528060200260200182016040528015610863578160200160208202803683370190505b50905060005b81518160ff1610156108be57858160ff1686018151811061088657fe5b6020026020010151828260ff168151811061089d57fe5b6001600160f81b031990921660209283029190910190910152600101610869565b50949350505050565b60f81c60011490565b6108d8610bb0565b6000825167ffffffffffffffff811180156108f257600080fd5b5060405190808252806020026020018201604052801561092c57816020015b610919610be4565b8152602001906001900390816109115790505b50905060005b83518160ff16101561097c5761095a610955858360ff168460010160ff1661081c565b610991565b828260ff168151811061096957fe5b6020908102919091010152600101610932565b50604080516020810190915290815292915050565b610999610be4565b600080836000815181106109a957fe5b602002602001015160f81c90506000806000806000600f8111156109c957fe5b60ff168560ff1614156109eb5760009550600193506000925060009150610abe565b600186600f8111156109f957fe5b60ff161415610a175760019550600293506000925060019150610abe565b600286600f811115610a2557fe5b60ff161415610a435760029550600393506000925060029150610abe565b600386600f811115610a5157fe5b60ff161415610a6d575060039450600091506001905080610abe565b600486600f811115610a7b57fe5b60ff161415610a97575060049450600091506002905080610abe565b600586600f811115610aa557fe5b60ff161415610abe575060059450600091506006905060035b60006040518060a0016040528088600f811115610ad757fe5b815260ff968716602082015292861660408401529385166060830152509216608090920191909152949350505050565b6040518060800160405280610b1a610b87565b8152602001610b27610bb0565b815260200160608152602001610b3b610bc3565b905290565b6040518060a00160405280610b53610bb0565b8152602001610b60610bb0565b8152602001610b6d610bb0565b8152602001610b7a610bb0565b8152602001606081525090565b604080516080810190915260008082526020820190815260006020820152606060409091015290565b6040518060200160405280606081525090565b60405180606001604052806060815260200160608152602001606081525090565b6040805160a081019091528060008152600060208201819052604082018190526060820181905260809091015290565b600060a08284031215610c25578081fd5b50919050565b60008060008060808587031215610c40578384fd5b843567ffffffffffffffff80821115610c57578586fd5b9086019060c08289031215610c6a578586fd5b90945060208601359080821115610c7f578485fd5b610c8b88838901610c14565b94506040870135915080821115610ca0578384fd5b50610cad87828801610c14565b949793965093946060013593505050565b6020808252601690820152754e756d626572206f66207061727469636970616e747360501b604082015260600190565b60208082526017908201527f5369676e6572206973206e6f74206e6578744163746f72000000000000000000604082015260600190565b6000808335601e19843603018112610d3b578283fd5b83018035915067ffffffffffffffff821115610d55578283fd5b6020908101925081023603821315610d6c57600080fd5b9250929050565b6000808335601e19843603018112610d89578283fd5b83018035915067ffffffffffffffff821115610da3578283fd5b602001915036819003821315610d6c57600080fdfea26469706673582212205ab4dfd9f51a28a15e78631a8a69e386ea6c58dcb16e76a36ea5012f65712d1864736f6c63430007060033"
