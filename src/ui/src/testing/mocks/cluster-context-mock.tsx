/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import * as React from 'react';

import { ClusterContext, ClusterContextProps } from 'app/common/cluster-context';
import { GQLClusterStatus } from 'app/types/schema';

export const CLUSTER_CONTEXT_DEFAULTS: ClusterContextProps = {
  loading: false,
  selectedClusterID: '',
  selectedClusterName: '',
  selectedClusterPrettyName: '',
  selectedClusterUID: '',
  selectedClusterVizierConfig: {
    passthroughEnabled: true,
  },
  selectedClusterStatus: GQLClusterStatus.CS_UNKNOWN,
  selectedClusterStatusMessage: '',
  setClusterByName: jest.fn(),
};

export const MockClusterContextProvider: React.FC = ({ children }) => (
  <ClusterContext.Provider value={CLUSTER_CONTEXT_DEFAULTS}>
    {children}
  </ClusterContext.Provider>
);
