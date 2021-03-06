package easter_egg

import (
	"hash/fnv"
	"math/rand"
	"strconv"
	"time"
)

var ASCII = []string{
	`
        ░░░░░░░░░░░█▀▀░░█░░░░░░
        ░░░░░░▄▀▀▀▀░░░░░█▄▄░░░░
        ░░░░░░█░█░░░░░░░░░░▐░░░
        ░░░░░░▐▐░░░░░░░░░▄░▐░░░
        ░░░░░░█░░░░░░░░▄▀▀░▐░░░
        ░░░░▄▀░░░░░░░░▐░▄▄▀░░░░
        ░░▄▀░░░▐░░░░░█▄▀░▐░░░░░
        ░░█░░░▐░░░░░░░░▄░█░░░░░
        ░░░█▄░░▀▄░░░░▄▀▐░█░░░░░
        ░░░█▐▀▀▀░▀▀▀▀░░▐░█░░░░░
        ░░▐█▐▄░░▀░░░░░░▐░█▄▄░░
        ░░░▀▀░▄TSM▄░░░▐▄▄▄▀░░░`,

	`
        ░░░░░░░░░░░░░░░░░░░░░░░
        ░▄▄▄░░▄▄▄▄░░░▄▄▄▄░░░▄▄░
        ▐░░▐▄▀░░░░▀▄▀░░░░▀▄▐░░▌
        ▐░░░▌░░░░░▄▀▀▄░░░░░▌░░▌
        ▐░▐░▐░░░░░▌░▌░▌░░░▐░▌░▌
        ░▀▀░░▌░░▌░▀▌▐▀░▐░░▌░▀▀░
        ░░░▌░▐░░▐▄▀▌▐▀▄▌░▐░░▐░░
        ░░░▐░░░░▐░░▀▀░░▌░░░░▌░░
        ░░░░▌░░░▌░░▐░░░▐░░░▐░░░
        ░░░░▐░░▄▐░▀░░▀░▌▄░░▌░░░
        ░░░░░▀▀░░▀███▀░▀▀░░░░░░`,
	`
                      _______ 
                  .-^^  ^.^  ^^-.
                .^      .^.      ^.
              /|         Q         |
             / |        :|:        |\
            / /         . .         \^\
           / /        ,/: :\,        \ ^\
          / /        / (_i_) \        \  ^\
         / /       /^   | |   ^\       \   |
        | /      /^     ^-^     ^\      \.^
        ||   / |/                 \ \ \  \
        || ,^-./                   \ |--. |
        | \    \                    /     |
        |\ \    |                  |     /
        \_^-\   \                  |    /
              \  \                 /   /
               \  \               /   /
              __\ \^,            /   /__
             /-  | | \          / / (  -\
             ~-._\    )        / | -^ _.-^
                  ^^^^        (   _.-^
                               ^^^`,
	`
        ░░░░░▄▀▓▓▄██████▄░░░ 
        ░░░░▄█▄█▀░░▄░▄░█▀░░░ 
        ░░░▄▀░██▄░░▀░▀░▀▄░░░ 
        ░░░▀▄░░▀░▄█▄▄░░▄█▄░░ 
        ░░░░░▀█▄▄░░▀▀▀█▀░░░░ 
        ░░▄▄▓▀▀░░░░░░░▒▒▒▀▀▀▓▄░ 
        ░▐▓▒░░▒▒▒▒▒▒▒▒▒░▒▒▒▒▒▒▓ 
        ░▐▓░█░░░░░░░░▄░░░░░░░░█░ 
        ░▐▓░█░░░(◐)░░▄█▄░░(◐)░░░█ 
        ░▐▓░░▀█▄▄▄▄█▀░▀█▄▄▄▄█▀░`,
	`
        ▒▒▒░░░░░░░░░░▄▐░░░░
        ▒░░░░░░▄▄▄░░▄██▄░░░
        ░░░░░░▐▀█▀▌░░░░▀█▄░
        ░░░░░░▐█▄█▌░░░░░░▀█▄
        ░░░░░░░▀▄▀░░░▄▄▄▄▄▀▀
        ░░░░░▄▄▄██▀▀▀▀░░░░░
        ░░░░█▀▄▄▄█░▀▀░░░░░░
        ░░░░▌░▄▄▄▐▌▀▀▀░░░░░
        ░▄░▐░░░▄▄░█░▀▀░░░░░
        ░▀█▌░░░▄░▀█▀░▀░░░░░
        ░░░░░░░░▄▄▐▌▄▄░░░░░
        ░░░░░░░░▀███▀█░▄░░░
        ░░░░░░░▐▌▀▄▀▄▀▐▄░░░
        ░░░░░░░▐▀░░░░░░▐▌░░
        ░░░░░░░█░░░░░░░░█░░
        ░░░░░░▐▌░░░░░░░░░█░`,
	`
        ░░░░░░░░░░▄▄▄▄░░░░░░ 
        ░░░░░░░▄▀▀▓▓▓▀█░░░░░ 
        ░░░░░▄▀▓▓▄██████▄░░░ 
        ░░░░▄█▄█▀░░▄░▄░█▀░░░ 
        ░░░▄▀░██▄░░▀░▀░▀▄░░░ 
        ░░░▀▄░░▀░▄█▄▄░░▄█▄░░ 
        ░░░░░▀█▄▄░░▀▀▀█▀░░░░ 
        ░░░░░░░█▄▄░░░░█░░░░░ 
        ░░░░░░░█░░░░▀▀█░░░░░ 
        ░░░░░░░█▀▀▀░▄▄█░░░░░ 
        ░░░░░░░█░░░░░░█▄░░░░ 
        ▄▄▄▄██▀▀░░░░░░░▀██░░ 
        ░▄█▀░▀░░░░▄░░░░░░█▄▄ 
        ▀▀█▄▄▄░░░▄██░░░░▄█░░`,
	`
        ░░░░░░░░░▓▓▓▓▀█░░░░░░░░░░░░░ 
        ░░░░░░▄▀▓▓▄██████▄ 
        ░░░░░▄█▄█▀░░▄░▄░█▀ 
        ░░░░▄▀░██▄░░▀░▀░▀▄ 
        ░░░░▀▄░░▀░▄█▄▄░░▄█▄ 
        ░░░░░░▀█▄▄░░▀▀▀█▀ 
        ░░░░░░█░░░░░░░░▄▀▀░▐ 
        ░░░░▄▀░░░░░░░░▐░▄▄▀ 
        ░░▄▀░░░▐░░░░░█▄▀░▐ 
        ░░█░░░▐░░░░░░░░▄░█ 
        ░░░█▄░░▀▄░░░░▄▀▐░█ 
        ░░░█▐▀▀▀░▀▀▀▀░░▐░█ 
        ░░▐█▐▄░░▀░░░░░░▐░█▄▄ 
        ░░░▀▀░OTHER LISTS░░ ▄▄▄▀`,
	`
        ▄▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
        ▓▓▓░▓▓▓▓░▄▓░▄▄▄▄▄▄░░▄▄▄▄▄▓▓▓ 
        ▓▓▓░▓▓░░▓▓▓░▓▓▓▓▓▓░▓▓▓▓▓▓▓▓▓ 
        ▓▓▓░▄▄░▄▓▓▓░▄▄▄▄▄▓░▓▓▓▓▓▓▓▓▓ 
        ▓▓▓░▓▓▓░░▓▓░▓▓▓▓▓▓░▓▓▓▓▓▓▓▓▓ 
        ▓▓▓▄▓▓▓▓▓▄▓▄▓▓▓▓▓▓▓▄▄▄▄▄▄▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░▓▓▓▓▓▓ 
        ▓▓▓▓░░░░░▓▓░░░░░░░░░░░░░░▓▓▓ 
        ▓▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓▓ 
        ▓▓▓▓░░░░░░░░░░░░░░░░░░░░▄▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░░▄▓▓▓▓ 
        ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄▄░░▄▄▓▓▓▓▓▓▓`,
	`
        ▒▒▒▒▒▄██████████▄▒▒▒▒▒ 
        ▒▒▒▄██████████████▄▒▒▒ 
        ▒▒██████████████████▒▒ 
        ▒▐███▀▀▀▀▀██▀▀▀▀▀███▌▒ 
        ▒███▒▒▌■▐▒▒▒▒▌■▐▒▒███▒ 
        ▒▐██▄▒▀▀▀▒▒▒▒▀▀▀▒▄██▌▒ 
        ▒▒▀████▒▄▄▒▒▄▄▒████▀▒▒ 
        ▒▒▐███▒▒▒▀▒▒▀▒▒▒███▌▒▒ 
        ▒▒███▒▒▒▒▒▒▒▒▒▒▒▒███▒▒ 
        ▒▒▒██▒▒▀▀▀▀▀▀▀▀▒▒██▒▒▒ 
        ▒▒▒▐██▄▒▒▒▒▒▒▒▒▄██▌▒▒▒ 
        ░░▄▄▓▀▀░░░░░░░▒▒▒▀▀▀▓▄░ 
        ░▐▓▒░░▒▒▒▒▒▒▒▒▒░▒▒▒▒▒▒▓ 
        ░▐▓░█░░░░░░░░▄░░░░░░░░█░ 
        ░▐▓░█░░░(◐)░░▄█▄░░(◐)░░░█ 
        ░▐▓░░▀█▄▄▄▄█▀░▀█▄▄▄▄█▀░`,
	`
        ▒▒▒▒▒▒▒▒▄▄▄▄▄▄▄▄▒▒▒▒▒▒▒▒ 
        ▒▒▒▒▒▄█▀▀░░░░░░▀▀█▄▒▒▒▒▒ 
        ▒▒▒▄█▀▄██▄░░░░░░░░▀█▄▒▒▒ 
        ▒▒█▀░▀░░▄▀░░░░▄▀▀▀▀░▀█▒▒ 
        ▒█▀░░░░███░░░░▄█▄░░░░▀█▒ 
        ▒█░░░░░░▀░░░░░▀█▀░░░░░█▒ 
        ▒█░░░░░░░░░░░░░░░░░░░░█▒ 
        ▒█░░██▄░░▀▀▀▀▄▄░░░░░░░█▒ 
        ▒▀█░█░█░░░▄▄▄▄▄░░░░░░█▀▒ 
        ▒▒▀█▀░▀▀▀▀░▄▄▄▀░░░░▄█▀▒▒ 
        ▒▒▒█░░░░░░▀█░░░░░▄█▀▒▒▒▒ 
        ▒▒▒█▄░░░░░▀█▄▄▄█▀▀▒▒▒▒▒▒ 
        ▒▒▒▒▀▀▀▀▀▀▀▒▒▒▒▒▒▒▒▒▒▒▒▒`,
	`
        ▓▓▓▓
        ▒▒▒▓▓
        ▒▒▒▒▒▓
        ▒▒▒▒▒▒▓
        ▒▒▒▒▒▒▓
        ▒▒▒▒▒▒▒▓
        ▒▒▒▒▒▒▒▓▓▓
        ▒▓▓▓▓▓▓░░░▓
        ▒▓░░░░▓░░░░▓
        ▓░░░░░░▓░▓░▓
        ▓░░░░░░▓░░░▓
        ▓░░▓░░░▓▓▓▓
        ▒▓░░░░▓▒▒▒▒▓
        ▒▒▓▓▓▓▒▒▒▒▒▓
        ▒▒▒▒▒▒▒▒▓▓▓▓
        ▒▒▒▒▒▓▓▓▒▒▒▒▓
        ▒▒▒▒▓▒▒▒▒▒▒▒▒▓
        ▒▒▒▓▒▒▒▒▒▒▒▒▒▓
        ▒▒▓▒▒▒▒▒▒▒▒▒▒▒▓
        ▒▓▒▓▒▒▒▒▒▒▒▒▒▓
        ▒▓▒▓▓▓▓▓▓▓▓▓▓
        ▒▓▒▒▒▒▒▒▒▓
        ▒▒▓▒▒▒▒▒▓`,
	`
             ROFL:ROFL:ROFL:ROFL
                  ___ ^_____
         L    ___/         [ ]
        LOL===_
         L     \_____________]
                 ___I______I__`,
	`
               ▄▀▀▀▀▀▀▀▀▀▀▄▄
            ▄▀▀             ▀▄
          ▄▀                  ▀▄
          █                     ▀▄
         ▐▌        ▄▄▄▄▄▄▄       ▐▌
         █           ▄▄▄▄  ▀▀▀▀▀  █
        ▐▌       ▀▀▀▀     ▀▀▀▀▀   ▐▌
        █         ▄▄▀▀▀▀▀    ▀▀▀▀▄ █
        █                ▀   ▐     ▐▌
        ▐▌         ▐▀▀██▄      ▄▄▄ ▐▌
         █           ▀▀▀      ▀▀██  █
         ▐▌    ▄             ▌      █
          ▐▌  ▐              ▀▄     █
           █   ▌        ▐▀    ▄▀   ▐▌
           ▐▌  ▀▄        ▀ ▀ ▀▀   ▄▀
           ▐▌  ▐▀▄                █
           ▐▌   ▌ ▀▄    ▀▀▀▀▀▀   █
           █   ▀    ▀▄          ▄▀
          ▐▌          ▀▄      ▄▀
         ▄▀   ▄▀        ▀▀▀▀█▀
        ▀   ▄▀          ▀   ▀▀▀▀▄▄▄▄▄`,
	`
        █                                                                             █
        █░▀▓█▀ ███▄    █   ██████ ▓█████ ▄████▄  █    ██  ██▀███   ██░████████▓▓██   ██
        ▓ ░██  ██ ▀█░  █ ▒██    ▒ ▓█   ▀▒██▀ ▀█  ██  ▓██▒▓██ ▒ ██▒▓██░▓  ██▒ ▓▒ ▒██  ██
        ▓ ░██ ▓██ ░▀█ ██▒░ ▓██▄   ▒███  ▒▓█    ▄▓██  ▒██░▓██ ░▄█ ▒▒██ ▒ ▓██░ ▒░  ▒██ ██
        ▒ ░█▓░▓██▒ ░▐▌██▒  ▒   ██▒▒██  ▄▒▓▓▄ ▄██▓▓█  ░██░▒██▀▀█▄  ░██░░ ▓██▓ ░   ░ ▐██▓
        ░░▄██▄▒██░   ▓██░▒██████▒▒░▒████▒ ▓███▀ ▒▒█████▓ ░██  ▒██▒░██░  ▒██▒ ░   ░ ██▒▓
        ░░ ░░░░ ▒    ▒     ▒▓▒ ▒     ▒░   ░▒ ▒  ░  ▒ ▒ ▒ ░ ▒  ░▒ ░░▓    ▒ ░░      ██▒░▓
        ░  ░                ░                                  ░                ▓██   ▒`,
}

func GetAscii(ip string) string {
	rand.Seed(FNV32a(ip + strconv.Itoa(time.Now().Day())))
	return ASCII[rand.Int()%len(ASCII)]
}

func FNV32a(text string) int64 {
	algorithm := fnv.New64a()
	_, _ = algorithm.Write([]byte(text))
	return int64(algorithm.Sum64())
}
