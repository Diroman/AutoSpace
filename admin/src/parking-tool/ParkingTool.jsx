import React from "react";
import {TwoDimensionalImage} from "../tool-src/apps/two-dimensional-image";
import {fetchData} from "../common/helpers/fetchData";

const ParkingTool = () => {

    const[data, setData] = React.useState(undefined);
    const[pictureUrl, setPictureUrl] = React.useState('');

    React.useEffect(() => {
        fetchData(`http://192.168.31.44:8080/get-frame`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            mode: 'no-cors',
        })
            .then(data => console.log(data))
            .catch(err => console.error(err));
    }, [])


    const handleSubmit = (r) => {
        console.log(r);
        setData(r);
        fetchData(`http://192.168.31.44:8080/set-polygon`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(r),
            mode: 'no-cors',
        })
            .then(data => console.log(data))
            .catch(err => console.error(err));
    }


    const options = {
        id: '0',
        value: 'root',
        children: [],
    };
    return (
        <div>
            <TwoDimensionalImage
                hasSkipButton
                onSkipClick={ handleSubmit }
                isDynamicOptionsEnable
                url='https://www.gtice.is/wp-content/uploads/2015/06/Snaefellsnes_Tour_Kirkjufell_by_KateI.jpg'
                imageWidth={ 800 }
                imageHeight={ 600 }
                options={ options }
                disabledOptionLevels={ [] }
            />
        </div>
    );
};

export default ParkingTool;
