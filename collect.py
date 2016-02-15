#!/usr/bin/env python

from lxml import html
import requests
import codecs
import collections

from providers import *

data_providers = [AllAccess(), FMQB(), BDSRadioCharts()]


def write_songs(songs, file):
    for song in songs:
        if song['title'] != '' and song['artist'] != '' :
            file.write("\"" + song['title'] + "\",\"" + song['artist'] + "\",\"" +
                       song['provider'] + "\"\n")


def normalize(possible_collection):
    if isinstance(possible_collection, collections.Sequence):
        if len(possible_collection) > 0:
            item = possible_collection[0]
        else:
            item = ""
    else:
        item = possible_collection

    return item


def get_songs(provider):
    page = requests.get(provider.url)
    tree = html.fromstring(page.content)

    songs_elements = tree.xpath(provider.item_root_xpath)

    songs = []
    for song_element in songs_elements:
        titles = song_element.xpath('.' + provider.title_child_xpath +
                                    '/text()')
        artists = song_element.xpath('.' + provider.artist_child_xpath +
                                     '/text()')

        title = normalize(titles)
        artist = normalize(artists)

        songs.append({ 'title': title.strip(),
                       'artist': artist.strip(),
                       'provider': provider.name.strip() })

    return songs


with codecs.open('./songs.csv', 'w+', 'utf-8') as file:
    for provider in data_providers:
        songs = get_songs(provider)
        write_songs(songs, file)
file.closed
