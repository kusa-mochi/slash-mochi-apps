interface Props {
    title: string | undefined
}

export default function TitleBarView({title}: Props) {
    return (
        <div className="w-screen h-12 bg-green-800 text-white shadow-md flex flex-row items-center justify-start px-4">
            <div className="text-2xl mr-4">Can I Use Up Food ?</div>
            <div>{title}</div>
        </div>
    )
}
