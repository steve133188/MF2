import {Pill} from "./Pill"
import {Badge} from "./Badge"
import {NormalButton} from "./Button";

export function CustomerProfileCategory() {
    return (
            <div className="customerProfileCategory">
                <div className="customerCategoryContainer">
                    <div className="categoryHeader">
                        <span className="categoryName">Assignee</span>
                        <NormalButton>+</NormalButton>
                    </div>
                    <div className="pillBadgeContainer">
                        <Pill color="lightYellow">MF</Pill><Pill color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill><Pill
                        color="lightYellow">MF</Pill>
                    </div>
                </div>
                <div className="customerCategoryContainer">
                    <div className="categoryHeader">
                        <span className="categoryName">Groups</span>
                        <NormalButton>+</NormalButton>
                    </div>
                    <div className="pillBadgeContainer"><Badge color="gp1">Group 1</Badge><Badge color="gp1">Group
                        1</Badge><Badge color="gp1">Group 1</Badge><Badge color="gp2">Group 2</Badge><Badge color="gp1">Group
                        1</Badge><Badge color="gp1">Group 1</Badge><Badge color="gp1">Group 1</Badge><Badge color="gp1">Group
                        1</Badge><Badge color="gp1">Group 1</Badge>
                    </div>
                </div>
            </div>
    )
}