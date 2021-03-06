// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
import bind from 'autobind-decorator'
import { connect } from 'react-redux'
import { defineMessages } from 'react-intl'
import { Redirect } from 'react-router-dom'
import { Container, Row, Col } from 'react-grid-system'

import { withEnv } from '../../../lib/components/env'
import Button from '../../../components/button'
import Message from '../../../lib/components/message'
import IntlHelmet from '../../../lib/components/intl-helmet'
import sharedMessages from '../../../lib/shared-messages'

import style from './login.styl'

const m = defineMessages({
  welcome: 'Welcome to {stackConsole}',
  login: 'You need to be logged in to use this site',
  loginWithStackAccount: 'Login with your TTN Stack Account',
})

@withEnv
@connect(state => ({
  user: state.user.user,
}))
@bind
export default class Login extends React.PureComponent {

  redirectToLogin () {
    const { env, location } = this.props
    const { from } = location.state || { from: env.app_root }

    window.location = `${env.app_root}/api/auth/login?path=${from}`
  }

  render () {
    const { user, env } = this.props

    // dont show the login page if the user is already logged in
    if (Boolean(user)) {
      return <Redirect to={env.app_root} />
    }

    return (
      <div className={style.login}>
        <IntlHelmet>
          <title><Message content={sharedMessages.login} /></title>
        </IntlHelmet>
        <Container>
          <Row>
            <Col>
              <Message
                className={style.loginHeader}
                values={{ stackConsole: 'TTN Stack Console' }}
                component="h2"
                content={m.welcome}
              />
              <Message className={style.loginSub} content={m.login} />
              <Button message={m.loginWithStackAccount} onClick={this.redirectToLogin} />
            </Col>
          </Row>
        </Container>
      </div>
    )
  }
}
