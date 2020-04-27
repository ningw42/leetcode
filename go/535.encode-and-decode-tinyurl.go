type Codec struct {
    storage map[string]string
}


func Constructor() Codec {
    return Codec{storage: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    for {
        h := md5.New()
        io.WriteString(h, longUrl)
        io.WriteString(h, strconv.Itoa(rand.Intn(10000)))
        short := string(h.Sum(nil)[:6])
        if _, duplicated := this.storage[short]; !duplicated {
            this.storage[short] = longUrl
            return short
        }
    }
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return this.storage[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
