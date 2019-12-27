import sys
from bs4 import BeautifulSoup
import requests

def show_scores(name):
    cur_scores = {}
    for i in range(25):
        p1, p2 = get_scores(i+1)
        for k, v in p1.iteritems():
            cur_scores[k] = cur_scores[k] + v if k in cur_scores else v
        for k, v in p2.iteritems():
            cur_scores[k] = cur_scores[k] + v if k in cur_scores else v
        my_score = cur_scores[name] if name in cur_scores else 0
        print('Day {}:\t{}\t#{}'.format(i+1, my_score, 1+len({k:v for k,v in cur_scores.iteritems() if v > my_score})))

# returns two dicts (for each part) where keys are names and values are points
def get_scores(day):
    URL = 'https://adventofcode.com/2019/leaderboard/day/{}'.format(day)
    page = requests.get(URL)
    doc = BeautifulSoup(page.content, 'html.parser')

    entries = doc.find_all('div', class_='leaderboard-entry')

    part2 = entries[:100]
    part1 = entries[100:]

    def rowToName(el):
        return el.find_all(text=True)[4]

    return {rowToName(part1[i]):100-i for i in range(100)}, {rowToName(part2[i]):100-i for i in range(100)},

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('Usage: python scrape_score.py <name>')
        sys.exit(1)
    show_scores(sys.argv[1])
