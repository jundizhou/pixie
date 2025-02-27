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

import {
  Button,
  Tab,
  Tabs,
  TableCell,
  Tooltip,
} from '@mui/material';
import { Theme } from '@mui/material/styles';
import { createStyles, makeStyles, withStyles } from '@mui/styles';
import { Link } from 'react-router-dom';

import { StatusGroup } from 'app/components';
import { GQLClusterStatus } from 'app/types/schema';

const tooltipStyles = makeStyles(() => createStyles({
  tooltip: {
    margin: 0,
  },
}));

const SEC_MS = 1000;
const MIN_MS = 60 * SEC_MS;
const HOURS_MS = 60 * MIN_MS;

export function convertHeartbeatMS(lastHeartbeatMs: number): string {
  const result = [];
  let time = lastHeartbeatMs;
  const hours = Math.floor(time / (HOURS_MS));
  if (hours > 0) {
    result.push(`${hours} hours`);
    time %= HOURS_MS;
  }
  const minutes = Math.floor(time / MIN_MS);
  if (minutes > 0) {
    result.push(`${minutes} min`);
    time %= MIN_MS;
  }
  const seconds = Math.floor(time / SEC_MS);
  result.push(`${seconds} sec`);
  return `${result.join(' ')} ago`;
}

export function agentStatusGroup(status: string): StatusGroup {
  if (['AGENT_STATE_HEALTHY'].indexOf(status) !== -1) {
    return 'healthy';
  } if (['AGENT_STATE_UNRESPONSIVE'].indexOf(status) !== -1) {
    return 'unhealthy';
  }
  return 'unknown';
}

export function clusterStatusGroup(status: GQLClusterStatus): StatusGroup {
  if (['CS_HEALTHY', 'CS_CONNECTED'].indexOf(status) !== -1) {
    return 'healthy';
  } if (['CS_UPDATING'].indexOf(status) !== -1) {
    return 'pending';
  } if (['CS_UNHEALTHY', 'CS_UPDATE_FAILED'].indexOf(status) !== -1) {
    return 'unhealthy';
  } if (['CS_DEGRADED'].indexOf(status) !== -1) {
    return 'degraded';
  }
  return 'unknown';
}

export function getClusterDetailsURL(clusterName: string): string {
  return `/admin/clusters/${clusterName}`;
}

export function podStatusGroup(status: string): StatusGroup {
  switch (status) {
    case 'RUNNING':
    case 'SUCCEEDED':
      return 'healthy';
    case 'FAILED':
      return 'unhealthy';
    case 'PENDING':
      return 'pending';
    case 'UNKNOWN':
    default:
      return 'unknown';
  }
}

export function containerStatusGroup(status: string): StatusGroup {
  switch (status) {
    case 'CONTAINER_STATE_RUNNING':
      return 'healthy';
    case 'CONTAINER_STATE_TERMINATED':
      return 'unhealthy';
    case 'CONTAINER_STATE_WAITING':
      return 'pending';
    case 'CONTAINER_STATE_UNKNOWN':
    default:
      return 'unknown';
  }
}

// eslint-disable-next-line react-memo/require-memo
export const AdminTooltip = ({
  children, title,
}: { children: React.ReactElement, title: string }): React.ReactElement => {
  const classes = tooltipStyles();
  return (
    <Tooltip title={title} placement='bottom-start' classes={classes} >
      {children}
    </Tooltip>
  );
};
AdminTooltip.displayName = 'AdminTooltip';

// eslint-disable-next-line react-memo/require-memo
export const StyledTabs = withStyles(() => createStyles({
  root: {
    flex: 1,
  },
  indicator: {
    height: '4px',
  },
}))(Tabs);

// eslint-disable-next-line react-memo/require-memo
export const StyledTab = withStyles((theme: Theme) => createStyles({
  root: {
    fontSize: theme.typography.subtitle1.fontSize,
    fontWeight: theme.typography.fontWeightLight,
    textTransform: 'none',
    minWidth: theme.spacing(20),
    '&:focus': {
      color: theme.palette.foreground.two,
    },
  },
}))(Tab);

// eslint-disable-next-line react-memo/require-memo
export const StyledTableHeaderCell = withStyles((theme: Theme) => createStyles({
  root: {
    fontWeight: theme.typography.fontWeightLight,
    fontSize: '14px',
    borderBottom: 'none',
  },
}))(TableCell);

// eslint-disable-next-line react-memo/require-memo
export const StyledTableCell = withStyles((theme: Theme) => createStyles({
  root: {
    fontWeight: theme.typography.fontWeightLight,
    fontSize: '14px',
    color: theme.palette.foreground.one,
    borderWidth: theme.spacing(1),
    borderColor: theme.palette.background.default,
  },
}))(TableCell);

// eslint-disable-next-line react-memo/require-memo
export const StyledLeftTableCell = withStyles((theme: Theme) => createStyles({
  root: {
    borderRadius: theme.shape.leftRoundedBorderRadius.large,
  },
}))(StyledTableCell);

// eslint-disable-next-line react-memo/require-memo
export const StyledRightTableCell = withStyles((theme: Theme) => createStyles({
  root: {
    borderRadius: theme.shape.rightRoundedBorderRadius.large,
  },
}))(StyledTableCell);

// eslint-disable-next-line react-memo/require-memo
export const LiveViewButton = withStyles((theme: Theme) => createStyles({
  root: {
    color: theme.palette.foreground.grey5,
    height: theme.spacing(4),
  },
// eslint-disable-next-line react-memo/require-memo
}))(({ classes }: any) => (
  <Button classes={classes} component={Link} to='/live'>
    Live View
  </Button>
));
