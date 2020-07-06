# Copyright Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from __future__ import print_function
import sys
import argparse
import yaml  # pyyaml


def equal_schema(args):
    kinds = args.kinds.split(",")
    versions = args.versions.split(",")
    with open(args.file, 'r') as stream:
        try:
            docs = yaml.safe_load_all(stream)
            for val in docs:
                if val is None:
                    continue
                kind = val["spec"]["names"]["kind"]
                if kind in kinds:
                    print("Checking schema equality in " + kind + "...")
                    for version in val["spec"]["versions"]:
                        if version["name"] in versions:
                            try:
                                schema
                            except NameError:
                                schema = version["schema"]["openAPIV3Schema"]
                            else:
                                if version["schema"]["openAPIV3Schema"] != schema:
                                    print(version["name"] + " of " +
                                          kind + " has a different schema")
                                    recursive_compare(
                                        version["schema"]["openAPIV3Schema"], schema)
                                    return -1
                    del schema
        except yaml.YAMLError as exc:
            print(exc)
            return -1
    return 0


def recursive_compare(d1, d2, level='openAPIV3Schema'):
    if isinstance(d1, dict) and isinstance(d2, dict):
        if d1.keys() != d2.keys():
            s1 = set(d1.keys())
            s2 = set(d2.keys())
            print('{:<20} + {} - {}'.format(level, s1-s2, s2-s1))
            common_keys = s1 & s2
        else:
            common_keys = set(d1.keys())

        for k in common_keys:
            recursive_compare(d1[k], d2[k], level='{}.{}'.format(level, k))

    elif isinstance(d1, list) and isinstance(d2, list):
        if len(d1) != len(d2):
            print('{:<20} len1={}; len2={}'.format(level, len(d1), len(d2)))
        common_len = min(len(d1), len(d2))

        for i in range(common_len):
            recursive_compare(d1[i], d2[i], level='{}[{}]'.format(level, i))

    else:
        if d1 != d2:
            print('{:<20} {} != {}'.format(level, d1, d2))


def get_parser():
    parser = argparse.ArgumentParser(
        description="Validate the generated CRDs")

    subparsers = parser.add_subparsers(title="actions")

    equal_schema_parser = subparsers.add_parser("check_equal_schema",
                                                add_help=False,
                                                description="Check if schemas of different versions within a Kind are equal")

    equal_schema_parser.add_argument("--kinds", dest="kinds",
                                     help="CRD Kinds to check")
    equal_schema_parser.add_argument("--versions", dest="versions",
                                     help="CRD Kind versions to check")
    equal_schema_parser.add_argument("--file", dest="file",
                                     help="CRD file to check")
    equal_schema_parser.set_defaults(func=equal_schema)

    return parser


if __name__ == "__main__":
    parser = get_parser()
    args = parser.parse_args()
    sys.exit(args.func(args))
