#!/bin/bash

# http://aatdownloads.getty.edu/
rsync -av --delete-after storage/example.xml root@0.0.0.0:/var/www/thesaurus/storage
