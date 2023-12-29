import Image from 'next/image'
import { Inter } from 'next/font/google'
import { OmikujiService } from '@/connect/omikuji_connect'
import { createConnectTransport } from '@connectrpc/connect-web'
import { PromiseClient, Transport, createPromiseClient } from '@connectrpc/connect'
import { useEffect, useState } from 'react'
import { OmikujiResponse_ResultLevel } from '@/connect/omikuji_pb'
import useSound from 'use-sound'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  var transport: Transport = createConnectTransport({baseUrl: 'http://os3-328-53953.vs.sakura.ne.jp:3090'})
  var connectClient: PromiseClient<typeof OmikujiService> = createPromiseClient(OmikujiService, transport)

  function ResultLevelToKanji(level: OmikujiResponse_ResultLevel): string {
    let ret: string = ""
    switch(level) {
      case OmikujiResponse_ResultLevel.DAIKICHI:
        ret = "大吉"
        break;
      case OmikujiResponse_ResultLevel.KICHI:
        ret = "吉"
        break;
      case OmikujiResponse_ResultLevel.CHUUKICHI:
        ret = "中吉"
        break;
      case OmikujiResponse_ResultLevel.SHOUKICHI:
        ret = "小吉"
        break;
      case OmikujiResponse_ResultLevel.HANNKICHI:
        ret = "半吉"
        break;
      case OmikujiResponse_ResultLevel.SUEKICHI:
        ret = "末吉"
        break;
        case OmikujiResponse_ResultLevel.SUESHOUKICHI:
        ret = "末小吉"
        break;
      case OmikujiResponse_ResultLevel.TAIRA:
        ret = "平"
        break;
      case OmikujiResponse_ResultLevel.KYOU:
        ret = "凶"
        break;
      case OmikujiResponse_ResultLevel.SHOUKYOU:
        ret = "小凶"
        break;
      case OmikujiResponse_ResultLevel.HANNKYOU:
        ret = "半凶"
        break;
      case OmikujiResponse_ResultLevel.SUEKYOU:
        ret = "末凶"
        break;
      case OmikujiResponse_ResultLevel.DAIKYOU:
        ret = "大凶"
        break;
    }

    return ret
  }

  const [omikujiResult, setOmikujiResult] = useState(ResultLevelToKanji(OmikujiResponse_ResultLevel.TAIRA))
  const [isShrineAnimating, setIsShrineAnimating] = useState(false)
  const [garagaraUseSound] = useSound('./garagara2.mp3')
  const [osaisenUseSound] = useSound('./osaisen3.mp3')

  function Omikuji(): void {
    console.log("opening omikuji...")

    setIsShrineAnimating((prev) => true)
    osaisenUseSound()
    setTimeout(() => garagaraUseSound(), 1600)

    connectClient.openOmikuji({}).then((res) => {
      setOmikujiResult(ResultLevelToKanji(res.result))
    }).catch((reason) => {
      console.log(reason)
    })
  }

  return (
    // <main
    //   className={`flex min-h-screen flex-col items-center justify-between p-24 ${inter.className}`}
    // >
    //   <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
    //     <p className="fixed left-0 top-0 flex w-full justify-center border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:static lg:w-auto  lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
    //       Get started by editing&nbsp;
    //       <code className="font-mono font-bold">pages/index.tsx</code>
    //     </p>
    //     <p><button onClick={Omikuji}>おみくじを引く！</button></p>
    //   </div>
    // </main>
    <div className='flex flex-col items-center'>
      <div className={`shrine-layout animated ${isShrineAnimating ? "shrine" : ""}`}>
        <img src='jinja.png'/>
        <div className={`result-container result-animated ${isShrineAnimating ? "result--showing" : "result--not-showing"}`}>
          {omikujiResult}
        </div>
      </div>
      <div>
        <button onClick={Omikuji} className='open-button' hidden={isShrineAnimating}>おみくじを引く</button>
      </div>
      <div className='sound-credit-container'>
        <a href="https://pocket-se.info/" target="_blank">ポケットサウンド/効果音素材</a>
      </div>
    </div>
  )
}
