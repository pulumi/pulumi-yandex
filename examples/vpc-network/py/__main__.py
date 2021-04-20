"""A Python Pulumi program"""

import pulumi
import pulumi_yandex as yandex

default = yandex.VpcNetwork("default")

pulumi.export("network-id", default.id)