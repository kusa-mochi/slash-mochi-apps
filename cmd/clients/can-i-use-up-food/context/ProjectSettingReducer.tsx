import { DateType } from "@/types/Date"
import { ProjectSettingType } from "@/types/ProjectSetting"
import { ScriptProps } from "next/script"
import { createContext, useReducer } from "react"

type ProjectSettingAction = 
| {
    type: "setName"
    payload: {
        projectName: string
    }
}
| {
    type: "setPeriod"
    payload: {
        startDate: DateType
        endDate: DateType
    }
}

const projectSettingReducer = (state: ProjectSettingType, action: ProjectSettingAction): ProjectSettingType => {
    const next: ProjectSettingType = {...state}

    switch (action.type) {
        case "setName":
            next.projectName = action.payload.projectName
            break
        case "setPeriod":
            next.startDate = action.payload.startDate
            next.endDate = action.payload.endDate
            break
    }

    return next
}

function useProjectSettingReducer() {
    const [state, dispatch] = useReducer(projectSettingReducer, {
        projectName: "テストプロジェクト",
        startDate: {
            year: 2024,
            month: 6,
            date: 25,
        },
        endDate: {
            year: 2024,
            month: 7,
            date: 10,
        },
    })

    return { state, dispatch }
}

const defaultProjectSetting: ReturnType<typeof useProjectSettingReducer> = {
    state: {
        projectName: "テストプロジェクト",
        startDate: {
            year: 2024,
            month: 6,
            date: 25,
        },
        endDate: {
            year: 2024,
            month: 7,
            date: 10,
        },
    },
    dispatch: () => {},
}

const ProjectSettingContext = createContext(defaultProjectSetting)
function ProjectSettingProvider (props: ScriptProps) {
    return (
        <ProjectSettingContext.Provider value={useProjectSettingReducer()}>
            {props.children}
        </ProjectSettingContext.Provider>
    )
}

export { useProjectSettingReducer, defaultProjectSetting, ProjectSettingContext, ProjectSettingProvider }
