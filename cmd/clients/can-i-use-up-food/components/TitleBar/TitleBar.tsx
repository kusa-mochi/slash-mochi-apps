import { ReactNode } from "react";
import TitleBarView from "./TitleBarView";

interface Props {
    children: ReactNode
}

export default function TitleBar({children}: Props) {
    return (
        <TitleBarView title={children?.toString()}></TitleBarView>
    )
}
