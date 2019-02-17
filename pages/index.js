import './ethsplain.css'
import React, { useEffect } from 'react'
import cuid from 'cuid'
import fetch from 'isomorphic-fetch'

const sampleResponse = {
  'tokens': [
    {
      hex: 'f9aa01',
      text: 'Nonce',
      more: 'Long explanation about nonces'
    },
    {
      hex: '85012a05f200',
      text: 'GasPrice',
      more: 'Long gas price explantion'
    }, {
      hex: '8327c50e',
      text: 'Gas Limit',
      more: 'Long explanation about noncesnthuntohnuh nthunouhtn nthountohun t\n ntohuntouh'
    },
    {
      hex: '35fb136cbadbc168910b66a9f7c40b03e4bd467f',
      text: 'Destination Address',
      more: 'How addresses are derived. \n 1. Concatenate x + y coordinate of pubkey\n2. keccak256(that point)\n3. last 20 bytes of hash'
    }, {
      hex: '68910b66a9f7c40',
      text: 'Value',
      more: 'Cash money business'
    }, {
      hex: '3e4bd467f80b8441e9a695000000000000000000000000035fb136cbadbc168910b66a9f7c40b03e4bd467f000000000000000000000',
      text: 'Contract Data',
      more: 'Extra stuff if i detect er20 20. + how to hash function prototypes'
    }, {
      hex: 'f9',
      text: 'Signature V',
      more: 'Long explanation about nonces'
    }, {
      hex: 'aca0026a00320143282b77654f3eedf2c6d384346a4be52c902f66032',
      text: 'Signature S',
      more: 'Long explanation about nonces'
    }, {
      hex: 'aca0026a00320143282b77654f3eedf2c6d384346a4be52c902f66032',
      text: 'Signature V',
      more: 'Long explanation about nonces'
    }
  ]
}

const Home = ({ response = sampleResponse }) => {
  console.log('Home response', response)
  useEffect(() => {
    process.browser && window.init()
  }, [])

  return (
    <>
      <svg id='canvas' />
      <div id='command'>
        { response.Tokens.map((token, i) => <span key={cuid()} className='command0' helpref={`help-${i}`}><a>{token.Hex}</a></span>)}
      </div>
      <div style={{ height: response.Tokens.length * 10 + 'px' }} />
      <div id='help'>
        {response.Tokens.map(({ Text, More }, i) => (
          <pre key={cuid()} id={`help-${i}`} className='help-box help-synopsis'>{Text}<br />{More}</pre>
        ))}
      </div>
    </>
  )
}

Home.getInitialProps = async ({ req, query: { tx, verbose } }) => {
  try {
    console.log(tx)
    console.log("verbose", verbose)
  const res = await fetch(`http://localhost:8080/${tx}?verbose=${verbose}`)
  const json = await res.text()
  console.log(json)
    return { response: JSON.parse(json) }
  } catch (e) {
    console.log(e)
    return {}
  }
  return { }
}


export default Home

