import Image from 'next/image'
import { Inter } from 'next/font/google'
import { OmikujiService } from '@/connect/omikuji_connect'
import { createConnectTransport } from '@connectrpc/connect-web'
import { PromiseClient, Transport, createPromiseClient } from '@connectrpc/connect'
import { useEffect } from 'react'
import { OmikujiResponse_ResultLevel } from '@/connect/omikuji_pb'
// import { useClient } from '@/connect_client/useClient'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  var transport: Transport
  var connectClient: PromiseClient<typeof OmikujiService>
  useEffect(() => {
    const baseUrl: string = `http://${window.location.hostname}:3090`
    console.log(baseUrl)
    transport = createConnectTransport({
      baseUrl: baseUrl,
    })
    connectClient = createPromiseClient(OmikujiService, transport)
  }, [])

  function Omikuji(): void {
    console.log("opening omikuji...")
    connectClient.openOmikuji({}).then((res) => {
      alert(OmikujiResponse_ResultLevel[res.result])
    }).catch((reason) => {
      console.log("error")
    })
  }

  return (
    <main
      className={`flex min-h-screen flex-col items-center justify-between p-24 ${inter.className}`}
    >
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <p className="fixed left-0 top-0 flex w-full justify-center border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:static lg:w-auto  lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          Get started by editing&nbsp;
          <code className="font-mono font-bold">pages/index.tsx</code>
        </p>
        <p><button onClick={Omikuji}>おみくじを引く！</button></p>
      </div>
    </main>
  )
}
