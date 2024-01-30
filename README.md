# Excel To Xml Converter
Excel 파일의 최상단 Row는 아래 표와 같이 여러 Level, Input, attribute(Unit, Type, Format..)으로 구성되어 있다.   
* example_excel의 MyExampleData.xlsx 참고
    |level0|level1|...|level7|INPUT|Unit|Type|Format|
    |---|---|---|---|---|---|---|---|
    |General Information|Basic Information|...|Basic Information|Home|Blank| string | Blank |

최종 결과물은 아래와 같이 사용자는 Excel에 Level에 맞는 tag 값을 지정 가능해야되며 INPUT과 속성 tag인 Unit, Type, Format 값을 지정하여 해당 소프트웨어를 사용하여 결과물로 excel의 파일로 부터 xml형식의 파일을 받을 수 있다.

```
<General_Information>
    <Basic_Information>
        <Building_Name type="string">Home</Building_Name>
        <Region>
            <Do_Si type="string">a</Do_Si>
            <Si_Gun_Gu type="string">b</Si_Gun_Gu>
        </Region>
        <Usage type="string">c</Usage>
        <Year_of_Completion>
            <Year type="Integer" unit="year">2024</Year>
            <Month type="Integer" unit="month">9</Month>
        </Year_of_Completion>
        <Height>
            <Floor_Height type="Integer" unit="mm">3600</Floor_Height>
            <Ceiling_Height type="Integer" unit="mm">2600</Ceiling_Height>
        </Height>
        <Renovation_History>
            <Date type="Date" format="yyyy-mm-dd">8/16/2011</Date>
            <Renovation_Item type="string">change</Renovation_Item>
            <Attachment>Y</Attachment>
        </Renovation_History>
        <Other_Details>
            <Exterior_wall_with_a_distance_of_not_more_than_1.5_m_to_the_adjacent_land_boundary>
                <Presence_Confirmation type="string">O</Presence_Confirmation>
                <Existence_of_Exterior_Wall_Windows type="string">O</Existence_of_Exterior_Wall_Windows>
            </Exterior_wall_with_a_distance_of_not_more_than_1.5_m_to_the_adjacent_land_boundary>
            <A_house_within_2m_of_the_exterior_wall_of_the_building>
                <Presence_Confirmation type="string">X</Presence_Confirmation>
                <Existence_of_Exterior_Wall_Windows type="string">X</Existence_of_Exterior_Wall_Windows>
            </A_house_within_2m_of_the_exterior_wall_of_the_building>
            <Circular_Window>
                <Presence_Confirmation type="string">O</Presence_Confirmation>
            </Circular_Window>
            <Curtain_Wall_Spandrel>
                <Presence_Confirmation type="string">O</Presence_Confirmation>
            </Curtain_Wall_Spandrel>
            <Supply_power type="string">N</Supply_power>
        </Other_Details>
    </Basic_Information>
    <No_Basic_Information>
        <Exteriors>
            <Front>d</Front>
            <Right_Side>e</Right_Side>
            <Back>f</Back>
            <Left_Side>g</Left_Side>
        </Exteriors>
        <Interiors>
            <Room_Name type="string">load</Room_Name>
            <Photograph>Y</Photograph>
        </Interiors>
    </No_Basic_Information>
</General_Information>
```
## Install & Execute
    1. Install golang (go version: go version go1.19.5 linux/amd64)
    2. git clone
    3.  make excute file
        3.1. windows : make windows
        3.2. Linux : make Linux
    4. Excute your binary file
    5. Enter Your Excel File PATH (example: /home/keti/go/src/xmlConverter/example_excel or ./example_excel)
    6. Select Your Excel Number (example: 1. my_first_excel, 2. my_second_excel -> 2) And Enter
    7. Select Your Excel Sheet (example: 0.Sheet1, 1.Sheet2, 2.Sheet3 -> 0) And Enter
    8. Check Created xmlResult.txt 

### Replace Character
Level의 Tag값에 xml 형식에 유효하지 않는 character가 포함되어 있는 경우 해당 character를 아래의 표와 같이 변환한다.
추후 추가 되어야될 character가 있을 경우 추가해야된다.  
| Before character | After character | 
|---|---|
|  (blank) | _ (under bar) |
| / | _ |
| ( | _ |
| ) | _ |
| & | _ |
| + | _ |
| % | _ |
| ° | _ |
| – | _ |

### Warning ( Develop 해야되는 부분 )
* 현재는 Level7으로 제한을 두었지만 Interface 형식으로 구성하였기 때문에 Levelx 까지 변경 가능 (코드 수정 필요)
* 속성 tag는 사용자마다 다를수 있으니 Input 값 뒤의 속성 tag를 상관없이 추가 가능하도록 구현해야됨 (코드 수정 필요)

### xml Validator
결과물로 나온 txt파일 확인 가능 Web Site
https://elmah.io/tools/xml-formatter/