from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

import os
from ansible_collections.whizz.embedded.lib.bincall import BinaryModule

class ActionModule(BinaryModule):
    NAME = os.path.basename(__file__).split(".")[0].lower()

