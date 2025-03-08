# Convert AGE encryption keys to passphrases
A wrapper to sit on top of [AGE](https://github.com/FiloSottile/age) - A random tool I use infrequently.

---
# Key point
This doesn't remove or change anything to with `age-keygen` - I still let people way smarter than me do encryption

---
# Working with it
Note: I am not really a programmer, more of a hobbyist. I like the idea of taking the secret keys and encrypting
them like most password managers use, cryptocurrency wallets work or just other random things I don't know about
recovery phrases

I started with the [Bitcoin list](https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt) which works out 
to be exactly 11 bytes with 2048 entries.
I've got some random copy of a file called dict.txt on my pc which works out to be about 67k entries, so I could have
done at 16 bits and saved myself so much time, but I had already started trying to divide ^2 numbers by 11 so. 
Go figure.

I did try and add padding for a long time, before giving up on it and just assuming all text input and -1 char on the
final word, all the theory made sense in my head but getting it in practice was not happening. So it's a bit hacky

---
# Usage
Compile simplez with `go run *.go`. This will work with any string technically given they're all divisible by 8. But
the guy who made this hardcoded `AGE-SECRET-KEY-` to the start of the strings... Hint, you could ditch that.

---
# Example output
Some of the raw binary/numbers are removed because it's just ASCII/binary, nothin' special.
```
Encoded: AGE-SECRET-KEY-1YRV7SRYNV4US80G2LFU8TX5Z3FJYDP5D54UN4PV2C7HSQVKGKC2S4AU3CW
StringToBinary: 0011000101011.... You know the rest of this
Binary Chunks (11 bits): [00110001010 11001010100 10010101100 01101110101] ... you know the rest of this
Decimal Values: [394 1620 1196 885 425 357 458 1588 ] ... you know the rest of this
Words: [course skin night huge cry coach deer shoot fetch oil army mobile cave case female appear flag print mimic milk pitch peasant dolphin possible cry pencil fatal hamster air bomb art trouble fashion melt name model reflect odor error market festival old helmet]
Decoded String: AGE-SECRET-KEY-1YRV7SRYNV4US80G2LFU8TX5Z3FJYDP5D54UN4PV2C7HSQVKGKC2S4AU3CW
```


# License
None, do as you please, but I would like to see it added to the Age tool