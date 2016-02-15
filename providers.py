class AllAccess(object):
    def __init__(self):
        self.name = "All Access"
        self.url = "http://www.allaccess.com/alternative/future-releases"
        self.item_root_xpath = "//*[@id='body']/ul/li[1]/ul[2]/li/div/div//ul[contains(@class,'song-rows')]/li[contains(@class,'artist')]"
        self.title_child_xpath = "//b"
        self.artist_child_xpath = "//h5"



class FMQB(object):
    def __init__(self):
        self.name = "FMQB"
        self.url = "http://fmqb.com/specialty.asp"
        self.item_root_xpath = "//tr[6]/td/table/tbody/tr"
        self.title_child_xpath = "//td[3]"
        self.artist_child_xpath = "//td[2]/strong"


class BDSRadioCharts(object):
    def __init__(self):
        self.name = "BDS Radio Charts"
        self.url = "http://charts.bdsradio.com/bdsradiocharts/charts.aspx?formatid=7"
        self.item_root_xpath = "//*[@id='grdMostAdded']/table/tbody/tr"
        self.title_child_xpath = "//span[2]"
        self.artist_child_xpath = "//span[1]"
