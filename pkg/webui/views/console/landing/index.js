// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { Switch, Route } from 'react-router-dom'

import WithAuth from '../../../lib/components/with-auth'
import Overview from '../overview'
import Applications from '../applications'
import Gateways from '../gateways'
import Organizations from '../organizations'

import Breadcrumbs from '../../../components/breadcrumbs'

export default class Landing extends React.PureComponent {

  render () {
    const { path } = this.props.match
    return (
      <React.Fragment>
        <Breadcrumbs />
        <Switch>
          <WithAuth>
            <Route exact path={`${path}`} component={Overview} />
            <Route path={`${path}/applications`} component={Applications} />
            <Route path={`${path}/gateways`} component={Gateways} />
            <Route path={`${path}/organizations`} component={Organizations} />
          </WithAuth>
        </Switch>
      </React.Fragment>
      // TODO:  render not found
    )
  }
}
