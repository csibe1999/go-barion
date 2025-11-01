Bank card structure
This structure represents the details about a funding source that was used when completing a payment. At the moment this is only used when the funding source is a bank card.

Included in
Bank cards are used in the following structures:
FundingInformation

| Property name  | Property type |                    Description                    |
| :------------: | :-----------: | :-----------------------------------------------: |
|   MaskedPan    |    string     |     The last four digits of the card number.      |
|  BankCardType  |   CardType    |            The type of the bank card.             |
| ValidThruYear  |    string     | The 4-digit year part of the card validity date.  |
| ValidThruMonth |    string     | The 2-digit month part of the card validity date. |
