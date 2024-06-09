interface Props {
    title: string | undefined
}

export default function TitleBarView({title}: Props) {
    return (
        <>
            <div>title-bar-view</div>
            <div>{title}</div>
        </>
    )
}
