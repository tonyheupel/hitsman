# song-charts-collector

This script will parse pages housing song charts and create a file
called `songs.csv` that can be imported into Excel.  The columns are:

`Title`, `Artist`, `Site`

# usage

The scripts require python 2.7 or above.

# installation

Open your terminal / console application and go to this directory and enter:

```
$ pip install -r requirements.txt
```

This will attempt to install the prerequisite libraries used by the script.

**NOTE: There are sometimes issues with `libxml2` when trying to install `lxml`.
To resolve this on Mac OSX, follow the instructions to install the command-line
tools for XCode, run `$ xcode-select --install` on the command-line**

Once the installation completes, you can run the script:

```
$ ./collect.py
```

This will erase the existing `songs.csv` (if there is one) and write a new one.


# adding more providers or updating existing providers

Provider settings are configured in the `providers.py` file.
Providers are implemented as Python classes that support the following properties:

* `name` - A descriptive name of the provider (e.g., `All Access`)
* `url` - The url to scrape
* `item_root_xpath` - the XPath that indicates the "root" (list item, table, row, etc.) of each song in a table, collection, etc.
* `title_child_xpath` - the XPath WITHIN the item root HTML tree that contains the text of the title
* `artist_child_xpath` - the XPath WITHIN the item root HTML tree that contains the text of the artist's name

**NOTE: With 'child' XPath items, you would normally need to use `.` to indicate "within the current root element."  The script takes care of that and also of appending `/text()` to the end of the selector so that you can focus on just getting the core XPath for each item figured out.**

To add another one:

1. Using a tool like the `Scraper` Chrome extension, determine the XPaths required for a given url.
1. In the `providers.py` file, copy an existing class (from `class` through the `artist_child_xpath` line) and paste it to the end of the file
1. Rename the class on the top line (e.g., `class SampleClass(object):`)
1. Update the properties inside the `__init__(self)` method
1. Save the `proviers.py` file
1. Open the `collect.py` file and add the new class name to the `data_providers` line (making sure to construct a new instance by adding the open and close parens `()`).  For example, adding SampleClass:
```
data_providers = [AllAcess(), FMQB(), SampleClass()]
```

Save the `collect.py` file and run!


Enjoy!
