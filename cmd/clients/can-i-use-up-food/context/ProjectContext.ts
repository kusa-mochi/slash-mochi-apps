import { createContext } from "react"

export type ProjectContextItemType = {
    testChildName: string
}

export type ProjectContextType = {
    testName: string
    testNumber: number
    testChild: ProjectContextItemType
}

export const ProjectContext = createContext<ProjectContextType>({
    testName: "あばばば",
    testNumber: 123,
    testChild: {
        testChildName: "おべべべ",
    },
})
