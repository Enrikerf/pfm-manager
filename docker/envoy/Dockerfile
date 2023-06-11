# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM envoyproxy/envoy:v1.20.0

COPY docker/envoy/envoy.yaml /etc/envoy/envoy.yaml

WORKDIR /

RUN apt update -y && apt install -y build-essential curl wget && \
    wget https://github.com/ktr0731/evans/releases/download/v0.10.11/evans_linux_arm64.tar.gz && \
    tar -xzvf evans_linux_arm64.tar.gz && \
    rm evans_linux_arm64.tar.gz

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
