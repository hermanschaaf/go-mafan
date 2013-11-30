import codecs
import json
import math
from collections import defaultdict

f = codecs.open("data/dict.txt.big", "r", "utf-8")
w = codecs.open('data/dict.compressed.tab', 'w', 'utf8')

words = []

print "compressing dict..."

# zipf's law parameters:
N = 0  # total number of words
s = 1.0  # the value of the exponent characterizing the distribution.
q = 0.1

total_words = 0

for i, line in enumerate(f.readlines()):
    l = map(unicode.strip, line.split(' '))
    words.append((l[0], int(l[1]), l[2]))
    total_words += int(l[1])
    N += 1

words = sorted(words, key=lambda w: w[1], reverse=True)

print words[:10]

print "writing dictionary..."

normalizer = sum([1 / math.pow(n + q, s) for n in range(1, N+1)])
diff = 0

for k, word in enumerate(words):
    freq = (1 / math.pow(k+1 + q, s)) / normalizer * total_words
    w.write(u"{}\t{}\t{}\t{}\t{}\n".format(word[0], word[2], word[1], freq, freq-word[1]))
    diff += abs(freq-word[1])

    real_freq = word[1] * 1.0 / total_words

f.close()
w.close()

print "done."
print "Total estimation difference:", diff