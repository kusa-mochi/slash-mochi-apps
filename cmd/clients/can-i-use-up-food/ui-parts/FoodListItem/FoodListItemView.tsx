type Props = {
    name: string
}

export default function FoodListItemView(props: Props) {
    return (
        <>
            <div>foodlist-item-view</div>
            <div>{props.name}</div>
        </>
    )
}
