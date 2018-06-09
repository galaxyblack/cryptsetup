```
$ sudo ./example 
1
<nil> Got error while crypt_get_verity_info(): -22: false true false false false false false false
LUKS header information for /dev/loop0

Version:        1
Cipher name:    aes
Cipher mode:    xts-plain64
Hash spec:      sha256
Payload offset: 4096
MK bits:        256
MK digest:      34 c8 2f b0 45 76 83 76 47 f1 56 10 60 1a 60 b8 39 86 e3 39 
MK salt:        67 a6 53 7c dd b5 01 3a d2 c2 e9 04 01 14 cd e4 
                1a b1 9f d7 37 6a 3e 07 c4 66 9f 72 0b 60 b5 2f 
MK iterations:  243750
UUID:           85a8b378-4b45-4a04-8079-8629e8cf8e1a

Key Slot 0: DISABLED
Key Slot 1: ENABLED
        Iterations:             1914017
        Salt:                   5b 79 65 be f1 b1 33 97 21 b4 a7 9e 5a 6d 24 9d 
                                15 4c 5e 83 0b b6 ab 87 2c ab ae d0 08 72 71 7e 
        Key material offset:    264
        AF stripes:             4000
Key Slot 2: DISABLED
Key Slot 3: DISABLED
Key Slot 4: DISABLED
Key Slot 5: DISABLED
Key Slot 6: DISABLED
Key Slot 7: DISABLED
<nil>
[56 53 97 56 98 51 55 56 45 52 98 52 53 45 52 97 48 52 45 56 48 55 57 45 56 54 50 57 101 56 99 102 56 101 49 97]
8
&{1} <nil>
```
