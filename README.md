# ICRekt


Purely proof of concept, and designed to demonstrate how severe "minor" smart contract bugs can really get. This is *entirely* proof of concept and is in no way meant to be ran or executed. It allows the user to DOS the ICX contract, preventing transfers for as long as there is ethereum in the calling account to pay the gas.

The cryptocurrency space is far to plagued with amateur developers who people blindly, or willfully trust millions of dollars to. At the time of this commit ICX has a market cap of 777,254,618USD which isn't chump change. A company like this, promising revolutionary products and technology with an evaluation of 777 million USD should write unit tests, and bugs like this should be caught during testing. Smart contracts are immutable, so until such time as the contract is replaced it can be permanently DOS'd preventing anyone from transferring their tokens, effectively locking up 777 million USD. If a company can't catch a simple bug like this, how on earth can you now trust that the code they put out, the code that will run their blockchain isn't susceptible to a similar bug or something worse like the EOS remote code execution vulnerability?

Rumour has it that it's the second time the developer of this token made this mistake, if that really is true that is hugely dissapointing.

https://twitter.com/minhokim/status/1007800274252062720?s=19
https://twitter.com/rgrpark/status/1007807240256962561

# Usage

Make sure you replace the address with your infura API key, or whatever node you use. After which run a go build and you'll be good to spam.
