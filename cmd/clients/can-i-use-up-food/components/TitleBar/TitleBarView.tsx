import { ProjectContext } from "@/context/ProjectContext"
import { useContext } from "react"

interface Props {
    title: string | undefined
}

export default function TitleBarView({title}: Props) {
    const ctx = useContext(ProjectContext)
    return (
        <div className="w-screen h-12 bg-green-800 text-white shadow-md flex flex-row items-center justify-start px-4">
            <div className="text-2xl mr-4">Can I Use Up Food ?</div>
            <div>{title}|{ctx.testName}|{ctx.testNumber}|{ctx.testChild.testChildName}</div>
        </div>
    )
}
