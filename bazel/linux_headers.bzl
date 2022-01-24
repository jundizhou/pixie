# Copyright 2018- The Pixie Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_file")

# Linux headers are bundled in PEM image for PEM to setup the BPF runtime.
def linux_headers():
    # timeconst_<config_hz>.h headers are used by BCC runtime for compilation. But it's part of
    # generated headers not in the original linux source. So we pack them into separately.
    http_file(
        name = "timeconst_100",
        urls = ["https://storage.googleapis.com/pixie-dev-public/timeconst_100.h"],
        sha256 = "082496c45ab93af811732da56000caf5ffc9e6734ff633a2b348291f160ceb7e",
        downloaded_file_path = "timeconst_100.h",
    )

    http_file(
        name = "timeconst_250",
        urls = ["https://storage.googleapis.com/pixie-dev-public/timeconst_250.h"],
        sha256 = "0db01d74b846e39dca3612d96dee8b8f6addfaeb738cc4f5574086828487c2b9",
        downloaded_file_path = "timeconst_250.h",
    )

    http_file(
        name = "timeconst_300",
        urls = ["https://storage.googleapis.com/pixie-dev-public/timeconst_300.h"],
        sha256 = "91c6499df71695699a296b2fdcbb8c30e9bf35d024e048fa6d2305a8ac2af9ab",
        downloaded_file_path = "timeconst_300.h",
    )

    http_file(
        name = "timeconst_1000",
        urls = ["https://storage.googleapis.com/pixie-dev-public/timeconst_1000.h"],
        sha256 = "da0ba6765f2969482bf8eaf21249552557fe4d6831749d9cfe4c25f4661f8726",
        downloaded_file_path = "timeconst_1000.h",
    )

    http_file(
        name = "linux_headers_4_14_176_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.14.176-trimmed-pl3.tar.gz"],
        sha256 = "08d4daeba7b4c81454be8cdd3d3da65b3bcae8fedf0f3787690c85310bf91664",
        downloaded_file_path = "linux-headers-4.14.176.tar.gz",
    )

    http_file(
        name = "linux_headers_4_15_18_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.15.18-trimmed-pl3.tar.gz"],
        sha256 = "4d8ba1fdfa2f085a32457a1f2656c2c9cfc75fa5410d508579d4c6c06cf3ee2f",
        downloaded_file_path = "linux-headers-4.15.18.tar.gz",
    )

    http_file(
        name = "linux_headers_4_16_18_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.16.18-trimmed-pl3.tar.gz"],
        sha256 = "ff41877f051f67653801e10e546e5e29d98fde9cc2ec7981a7217ba853f5ab33",
        downloaded_file_path = "linux-headers-4.16.18.tar.gz",
    )

    http_file(
        name = "linux_headers_4_17_19_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.17.19-trimmed-pl3.tar.gz"],
        sha256 = "b4924b84ca3703d788eb7962b3df30fa13788978326d8f9f3a2913da6154b0a0",
        downloaded_file_path = "linux-headers-4.17.19.tar.gz",
    )

    http_file(
        name = "linux_headers_4_18_20_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.18.20-trimmed-pl3.tar.gz"],
        sha256 = "7260e6e7c18baaaa9f49913a1c756ab431b03ad8caa0ca15b34bfaedb232509a",
        downloaded_file_path = "linux-headers-4.18.20.tar.gz",
    )

    http_file(
        name = "linux_headers_4_19_118_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.19.118-trimmed-pl3.tar.gz"],
        sha256 = "bb41b959b34483f9206157e3195b7567c594d7595d3fdef94d579932684aed39",
        downloaded_file_path = "linux-headers-4.19.118.tar.gz",
    )

    http_file(
        name = "linux_headers_4_20_17_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-4.20.17-trimmed-pl3.tar.gz"],
        sha256 = "8024a42ead5dd474aa8190896f070df7c4c08c77228b1f18143779aa0ab877c5",
        downloaded_file_path = "linux-headers-4.20.17.tar.gz",
    )

    http_file(
        name = "linux_headers_3_10_0_tar_gz",
        urls = ["https://kdlibrary.oss-cn-beijing.aliyuncs.com/linux-headers-3.10.0-trimmed-pl3.tar.gz"],
        sha256 = "399829a43769b68e70c1b8396f14c71973ec262e9b8894388f4d208694428e21",
        downloaded_file_path = "linux-headers-3.10.0.tar.gz",
    )

    http_file(
        name = "linux_headers_5_0_21_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.0.21-trimmed-pl3.tar.gz"],
        sha256 = "3d044deec38127c05de3951cc5521370f500e8b85faed0497c447356dbd3af49",
        downloaded_file_path = "linux-headers-5.0.21.tar.gz",
    )

    http_file(
        name = "linux_headers_5_1_21_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.1.21-trimmed-pl3.tar.gz"],
        sha256 = "0f1669f5100c3f3f1ccc448401ba7ef88388ae08ede540b23df3744b9e81d344",
        downloaded_file_path = "linux-headers-5.1.21.tar.gz",
    )

    http_file(
        name = "linux_headers_5_2_21_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.2.21-trimmed-pl3.tar.gz"],
        sha256 = "31207779809d6e3b118ca6dd00bde382744d8ff05f53918658f21ac1bac6a423",
        downloaded_file_path = "linux-headers-5.2.21.tar.gz",
    )

    http_file(
        name = "linux_headers_5_3_18_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.3.18-trimmed-pl3.tar.gz"],
        sha256 = "ddac47a719244aee65f18e0d05d80f063b8d28db54258243590a27509d691753",
        downloaded_file_path = "linux-headers-5.3.18.tar.gz",
    )

    http_file(
        name = "linux_headers_5_4_35_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.4.35-trimmed-pl3.tar.gz"],
        sha256 = "b8eec62c5638642c4eb44f99cce5bf6862d3958e566733250c5e7aa4db2ccb13",
        downloaded_file_path = "linux-headers-5.4.35.tar.gz",
    )

    http_file(
        name = "linux_headers_5_5_19_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.5.19-trimmed-pl3.tar.gz"],
        sha256 = "3a8e2eded3645cbab6e495246616dbe0700bc9d39359cbf7859d402fcb339e34",
        downloaded_file_path = "linux-headers-5.5.19.tar.gz",
    )

    http_file(
        name = "linux_headers_5_6_19_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.6.19-trimmed-pl3.tar.gz"],
        sha256 = "ca3485b6f61696efbaf748f1edc9f307b9a0c7b208b545145277687f97e7e6f5",
        downloaded_file_path = "linux-headers-5.6.19.tar.gz",
    )

    http_file(
        name = "linux_headers_5_7_19_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.7.19-trimmed-pl3.tar.gz"],
        sha256 = "40a1484c3af86f7ef6e9f7bc497f2fdf6a9e052f0d80cff58f40f6dd4f53d158",
        downloaded_file_path = "linux-headers-5.7.19.tar.gz",
    )

    http_file(
        name = "linux_headers_5_8_18_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.8.18-trimmed-pl3.tar.gz"],
        sha256 = "fc7f5ac376e3af8b16f608dec7da5ce95e3dc78b87a172d545597e2800e5e7bf",
        downloaded_file_path = "linux-headers-5.8.18.tar.gz",
    )

    http_file(
        name = "linux_headers_5_9_16_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.9.16-trimmed-pl3.tar.gz"],
        sha256 = "2d68890188ff4b20c28704cda7e774afdd4729b15a2376f8677105d4851459f5",
        downloaded_file_path = "linux-headers-5.9.16.tar.gz",
    )

    http_file(
        name = "linux_headers_5_10_34_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.10.34-trimmed-pl3.tar.gz"],
        sha256 = "2dcf53abec0f49bffa5b5ceb65465c6d41ce620d304666955d37cd0f1d9e53f5",
        downloaded_file_path = "linux-headers-5.10.34.tar.gz",
    )

    http_file(
        name = "linux_headers_5_11_18_tar_gz",
        urls = ["https://storage.googleapis.com/pixie-dev-public/linux-headers-5.11.18-trimmed-pl3.tar.gz"],
        sha256 = "1ba2d5db17e2c913e8e0f0f99be35a83ffba89dd10cfe9d8051d9d8f3b7a75e4",
        downloaded_file_path = "linux-headers-5.11.18.tar.gz",
    )
