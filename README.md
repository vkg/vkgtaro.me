# vkgtaro.me

```
       _         _                                  
__   _| | ____ _| |_ __ _ _ __ ___   _ __ ___   ___ 
\ \ / / |/ / _' | __/ _' | '__/ _ \ | '_ ' _ \ / _ \
 \ V /|   < (_| | || (_| | | | (_) || | | | | |  __/
  \_/ |_|\_\__, |\__\__,_|_|  \___(_)_| |_| |_|\___|
            __/ |
           |___/           github.com/vkg/vkgtaro.me

Usage:

 $ curl -L vkgtaro.me
 $ curl -L vkgtaro.me/{SIZE}       (xs, s, m, l, xl)
 $ curl -L vkgtaro.me/{SIZE}/{OPT} (i, g)
 $ curl -L vkgtaro.me/h            (help)


Examples:

 $ curl -L vkgtaro.me/m            (medium)
 $ curl -L vkgtaro.me/l/i          (large, inverted)
 $ curl -L vkgtaro.me/xl/g         (extra large, grayscale)
```

[vkgtaro.me](https://vkgtaro.me) - curl vkgtaro in your terminal.

## How to use

From your terminal:

```sh
# Show vkgtaro
$ curl -L vkgtaro.me

# Show vkgtaro in any size you like
$ curl -L vkgtaro.me/{SIZE} (xs, s, m, l, xl)

# Show vkgtaro in any size/with option
$ curl -L vkgtaro.me/{SIZE}/{OPT} (i, g)

# Show how to see vkgtaro
$ curl -L vkgtaro.me/h
```

Example:

```sh
# Show medium sized vkgtaro
$ curl -L vkgtaro.me/m

# Show large sized inverted vkgtaro
$ curl -L vkgtaro.me/l/i

# Show extra large sized grayscale vkgtaro
$ curl -L vkgtaro.me/xl/g
```

From browser: access to [vkgtaro.me](https://vkgtaro.me)

## Development

* Generating vkgtaro ASCII art

Run `make gen`. It is depending on [im2a](https://github.com/tzvetkoff/im2a) command.

* Shipment

[vkgtaro.me](https://vkgtaro.me) is running on [vercel.com](https://vercel.com).

## Special Thanks

[vkgtaro.me](https://vkgtaro.me) is greatly inspired and highly respecting [dogs.sh](https://dogs.sh).
Thank you for amazing project!
