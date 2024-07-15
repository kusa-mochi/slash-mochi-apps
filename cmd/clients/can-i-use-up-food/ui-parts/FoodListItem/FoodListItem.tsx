import FoodListItemView from "./FoodListItemView";

type Props = {
    key: string
    name: string
}

export default function FoodListItem(props: Props) {
    return (
        <>
            <FoodListItemView
                name={props.name}
            ></FoodListItemView>
        </>
    )
}
