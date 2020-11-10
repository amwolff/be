package classifier

import (
	"encoding/json"
	"math"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	const f = `{"osCpu":{"duration":0},"languages":{"value":[["en-US"],["en-U` +
		`S"]],"duration":1},"colorDepth":{"value":24,"duration":0},"deviceMem` +
		`ory":{"value":8,"duration":0},"screenResolution":{"value":[812,375],` +
		`"duration":1},"availableScreenResolution":{"value":[812,375],"durati` +
		`on":0},"hardwareConcurrency":{"value":8,"duration":0},"timezoneOffse` +
		`t":{"value":480,"duration":0},"timezone":{"value":"America/Los_Angel` +
		`es","duration":24},"sessionStorage":{"value":true,"duration":1},"loc` +
		`alStorage":{"value":true,"duration":1},"indexedDB":{"value":true,"du` +
		`ration":0},"openDatabase":{"value":true,"duration":0},"cpuClass":{"d` +
		`uration":0},"platform":{"value":"iPhone","duration":0},"plugins":{"v` +
		`alue":[],"duration":0},"canvas":{"value":{"winding":true,"data":"dat` +
		`a:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPAAAACMCAYAAABCtSQoAAAgA` +
		`ElEQVR4Xu2dB5hcZbnHf9/M7G56IyRLElpCCSQkhBpCRwWliQhYaCGFBBTBci0YqQpyv` +
		`QiKSCCFUEQvoIJe4XoRMJBQQippBFJJSO9925z7/M+cszs7O7NzzvTZnfd59gnsfvU95` +
		`3/e+r2focDJwmoLHAscDfR3/j0I6AJ0dH4OcLaxBdjl/GwH1gEfRf8YzD53yxZWO+A4Y` +
		`IDzMxDoDGjONs6/7n93cPrtBvYDGkc/7n9rvgXAIuff+dFzYWV2H5iGfRT4IywtL4scM` +
		`FkcO6WhHVCdCXzO+TkeCKQ0WNNOYWADsAcQICszNG6TYfYC/2rP+r8cxe43zqL9mi/Q0` +
		`zqbgD1r+qR9zNUUzs90jNGUnsi6EctTwxbWyDxBwb3v6bI47xuysELAKQ5YPw8MBcrT3` +
		`Viu+9cCM4DXHUS9B1THLiIIDAHOAs4GzgE6ZWSlmupdZ2otYQbG1CUauQTgjPC8IAbJG` +
		`4AtrH7A5c7PqVCcX8dlwF+cn/fBv2g7AfgKcC1waJx3oi4MS9fDum2wvwbaV8ChB8Ih3` +
		`RPiE9D348/ACxjzaWzDEoALAnsZWUTOAWxhHQHcDlwDlDW3i2rC/CdzWcpOrucozqVXR` +
		`jadiUGWAvcBzwI1mRhQT+J0B8hfcyzx6UvgrzNgV73Z3jDTQV3hG6fD0c3yREvTEh/Am` +
		`CVu5xKAM/HACmOMpgAeNWEopu67GGOwLDlmZLVJWk5jwtjvp7psC0vK4zjgMq827T3Mo` +
		`jPlXEFf3mMjVdTxY95nOd8glKZZvIhtPMBc5rGFuVzBZJYwkqnM5woG0i3hNucAPwdeA` +
		`mSINk/yob0NHAXIlPdBp7wHXT6Ew5rpYwyMOBdO0TexWdJSteT7MWZmNICfsYbyI3M59` +
		`1l/ZbiRFt5yqeXbwKOf+ApYkzF8jifGzK5/lMOfbENZ1UQmjJXU9EUWlhxFvwGu8tXRN` +
		`oZf4naGcKmjW37CDv7OKr7HIL9DxW0/hSU8zHwbwLWEKWNiQgCvB24D/tv3zK+ArTn4A` +
		`fAK4LXITPKJD3a+AfFceaEg3HkF9FDDpCTn1ZR1Y8wN0d674xnHbdbrJQAnZV/hNWiQw` +
		`Fc+X06XbSswTOKJMXc0WeqoCecxcfQbXrdgYel1uwW4J1VXzaE8x6OcwcUc4nVaX+2e4` +
		`RMeZF4jAC/iKo6xI1QRkuh6BBBDdvoa3W3sF8DCmD4TMbO1dwJeCqjJ7RdNJ/eDUXLae` +
		`6NOY9pxESfyLNNsPaYEYG98K8RWDQC+8fEzbH3PMqcz4cZ34i72hskHEqp+CMPlWNatT` +
		`LhpAliGkeOHMemm6YyY1ItgzU87bS/rO/HvZ/e5e//sgRUEbe+U7NjD6MC3GcgrfMoMN` +
		`tKRcv6XCzk0JrYiaXg/c/klczmDSg6mPT/jBH7DAn7LAmoZbS9PKvVP+YCN7KOcAJvYz` +
		`0OcRl86cTezuItZPMd5/J5FXEAf/oPB/JgZrGY3PWnLMnaynr2NAHwDR/MZe1jFLj7HY` +
		`XzIKUyzd/AxsNwJO290DNYeju9ZER2psQpDK0IlaetK3GgAq5+cxCcBRyZ4HxTlejnqb` +
		`wpNLwYcp3IFcHIHuOogWPBJpF0oAA8Nh/JYZDvDbN4Fz78D3TvBll0wdzFYj3OGWcxzw` +
		`CWOBP4mH3CL+TpPcCYPWH/he7zGYnMQt1lXcTIruZ+XMKZ4I1AtW4UeNf46AuYp6oKHM` +
		`WnUqoRfm+GPVlIWWoMVOIGJoz9k1IQzMeG/0aaqkke+U3X6KS8///uZZ5w/KNyt81jep` +
		`oYwkzgbOaRO4yXbEfUrhlJHmGG8zCUcaoMzHvXhWcZzVr0EnskmW612AfxzZvMBm3iZC` +
		`+zu41nEEyxmFl+1QXgwf+DfXMIJdOcFltugfIt1vM7FBDA0p0JPYR8jeYEwcpArh0T5I` +
		`HIFyO8mW1FhV0k9AfYPwHVO7sdKQIrKCGdL0QDWd1G5KIltbJhF5EekXBK5p18EtkWxq` +
		`D90PB4G/SmS4qLvyy1fgoEHx39sv/tfOOxAuPhEsCwY8yvgT2Dm0NXOiBnHHY4Kvceqo` +
		`BcP8Kj5I9cgvzpcYn2Ll3iMoElu9ReilHLX1LIBPPqx4ZjAk3Yw44kxTUIPjR7MjY//E` +
		`8uaxYSxtzN6/G/BXN1rb9tbPvvDtWddyWtjXuALdnMBOAIs5WXAcP5NVypsKRnv/2Mff` +
		`jIAH8+LjOFYbrLfYqVd7aUXz7KMr1NO0AbwJq6juw0EOIbn+TYD+JadeEVcAL9vO7S68` +
		`bjdYqrjY1Z4OpokcdcCF8YBsFTfPzkAlkR0AVwFKPSjj0Bz9A/gs6gGCok3iSg7H5IaU` +
		`A6aYsrXDoHLTo4/8Lg/wZeGwOn6EAFjbgfrPTBvOu3Hca71Oq+Yd21O3cw3+YhK3uDXL` +
		`LR68UdO5ucmWitIsoUC/XPLBvCY8WdhmanNqtDugxn9xLWY8F1UVB/L/vJHO+wtCw/7R` +
		`8+v/nr7sG4vs9J2PIniAbgL5TzMsHoAR/+/XwBX8gwPMpSrHXV0H7W0YzLTbLdXxyYA7` +
		`sIUJnAWV9I3IYCP4wrm10tISaDNwEWAvmn60Ssul5asx+YAfIMDMgFY0lrtpf9qrOZI0` +
		`nZrTAN9cJSxucmR3sognd/QRhL48/3hL2dFYgaxNOF1aFsO15wJe6vgtseAh8G4itY4s` +
		`F5nqHmXV4FV9GGINY6ljLPV6dFmGv3suYubWjaAb3y8DIvlGJ7miTE/bfKoRo0fycSxk` +
		`+zf3/xoB2pDG7D43QlrDlz68L9O++EXav7niJsZYEs42aC5APAQ/sxo+qN5RWvZS+9mJ` +
		`HA//mjbwWMdiR2tQq8gTF8mAldEqbiSUAKebFZZi1cDSp/2IoGjASynmD5qck7J1dBc2` +
		`Edh23hZkbK3pXor1TtaQrtP6lCovAB+C1wZ8/R27IWH/wH9KiEchmm3gpkX1SgCYMy7t` +
		`n9f/u9L+TGf4yNWW1152kgxK35q2QDW8xkz/otYPIfhSh4fK29LhM65K8SRlS8zYWyD+` +
		`Bg9/rngzsAX9zw/cneFFTh4AM/ThQqm8+X6btmWwL9gju0Mc23gx1jERD5iJpfX28DRK` +
		`vRNvM0Cttk2sJxeLoBf5ArOI8zqRgCW2ipVWHauJOdfgeudLE/ZqHI2eZXAbhhpIaDon` +
		`DI1EmWLKlwrZ5df0kdMmSBOwE6Y07dGNPH1iANrqOM4u0MayHaod0g1AFjN1eoG63TuN` +
		`t9gMk/zTTtJtPip5QNYz2jkY0MIBn6EZenUzQ6M6YBlZ0+tYMLYUe5jvGTYqzcfMLPtf` +
		`z1ZfbZO69hJEfI432Yf7tGRnG2MYCoWFr/mNDvF8FbeoQ1BHuQ024P8Xd6hHSEmcjb9o` +
		`0I36j+BxXyPd7mQQ7iGI21nV6wTS2P8hBlspcoeZy177LH70YlfMY8f8j73cTKjOca2g` +
		`7ewn9G8xUdsr9cSpiNF4ix22BkTOhsgu1U6qaTg4bblHEmQ/Kej2h7oAHq1k8wsz7MyF` +
		`+XskvwSuAVSJTzLRSQpLgl8HiCW6kPQ25HE8fRdzTsNkDPMC0lDkA9A6eRRXmiZ2jJbh` +
		`etX58BLH8QMpvEfBOTYUi6nvOz/BKN9QQ/KCVl3sYw7aWMykmvmZTNZbdM6AOyBhRaWP` +
		`uHyUOU0t1HS9gJeYRvDPawyeRMFheRekzuq8EhAXuOciJSnWzawjkxIcssOlx2sD4EeQ` +
		`QJprmYPq/lUOHdAQ/70jfpw3Gu7FzECbjzqSAfrfOaaPyd1uxUe7+KvqARgWxZZemME3` +
		`ognKAekWO1eam31V5L1bS5Ne1aBVuAViFs2VUO/52D29dDJCfuPkWYgT/S9YGRTR5M81` +
		`UvAUuLcG/Qzm+1gknvguph51eoBbGHps68gaMRrlCNawnZu5C1bTf4ZJzKMnmnNrFdWg` +
		`SxZpK2D5kHlRhjXDjrVwdtyyP0LjM5SxdKdYCkU9i4YmQYR5VzGhR5+MVMJwFh/sxN3i` +
		`pwkv/9e5HtIaflKgBYSf+v/EJpck3KvFTO1agBbWMrlf6iYH6DWLpPwu8W+iXTWLzF6j` +
		`gEVJfJJevh6CYqVWi2ALawTHdW52fO7hf5g5R+W6twyfKppcDtgIk5xn14MPfzpSsVOY` +
		`+p8dm2VALawZHDKjxGvXkQ+n4evuRW1VaAncZK3r+GKv7G06AeA//C9FcWfhmCMzoqXK` +
		`M8cSGoMWVhTnAyGPC81vekVeHoqvSFaZm9VJ2iad5dsr09hTGZieclmKv29WQ40C2ALS` +
		`2XX3Iz3omXlv4Fzi3b1OVj4M06BI39TnYepPw3hr2epdcY4kBDAFpbyB5Uw6xxhydicO` +
		`R1IZ4BU0KK+IFROZy+SyVQt83+AL/par1g6GGPE4hLliQPNAfgu4M48rStj02oTd2dst` +
		`BY8kLK2dHpSQV/v9BOM+aX35qWWmeZAXABbWJK6kr6SwkVLERGhyh0l8sQBpWwrd8O7z` +
		`qWcmAEYJ4Ha0yS5aXT1e3SqqKUDFbQJh5COocIltYEa9td2YPfk/vaxrqKnRAAuOa6K/` +
		`tGmuAHFGnRasqEsWLKBCsehZWFGvU8PyuxUveZDnjXUVLVl/TOD7YPORVsnqAmALSydJ` +
		`tP9PvVHW6rYw262UMVu6uwiOXUECFJOW9rQiY4caP9/IZGqRemMjtL//ZGypHWIXz31D` +
		`uh0lf5fAShVfkxawtXfdL5bK4qjE1A6yODWltT/6/c6h52o1paPic53Dl956yJGSQonO` +
		`hXBde9zQHmIw2oD7JwyBKeQl7fBvba6cgHlnavscxfuIUpPXWst9pTDsidOym16gLPe4` +
		`+osrCdPso+vpUTxADxeJ4M1mkWYzaxkb309JkPQOW4nILsfrgAhDqQvbQooW3YsOGVx/` +
		`PDFBYLYohM7ArCOFLYyAItlv3DK73tj3+MYI5bHpWwD+MaZlFVD/5BJ7UoeK0RVsIYlu` +
		`QRxVgBsYSnBThnubXWOdwNLkPSVdO1EJR04gKCjmejv+9nFDtbZkllnaHtwBG0zdNmPt` +
		`/cmfiuVU9enOM59BkmGVckcaVQ68xtbytbVspKGztNZuoe+OZDAWoWOGaukgwKJyUmsP` +
		`gJj4p7MzDKAzahZdqVAX5I3dkuSxFNOsisX5oSyBeAf29X77RqIa9jJBhu8PTmK8gT8E` +
		`ZA3sZx9bCeEKhoOwOT5miO5RX+S0mOIB46UBspipxwBWDvQGUId2fJ2+Eu3eSjVvAllE` +
		`8DDZ1AZCtr2RNpUFWL1M4NTKofie+5sAVjui8F11PIZ820VuhuH2DZuc6T2W/mUdnShH` +
		`V3Zwkr2sDVu3z1sY7NzCvcgjrXt6GhyPxxd6EVnDrJV+D1soTuHU0EHtvOZLfllh5fRB` +
		`rVrS2e78sd2uyrWNs6nhkW2qq8aUjq+nExquipy7C5l16u+c3MqtErvKFFTR/D035pLq` +
		`rfs5Xi+FN26oKJ1Us21RtW30mF9/b/rOZJZKWG2w7HF1U5/015UMaM5G1g2uvqqDK3MH` +
		`IlS9/B/vKCCnLHSOlQ0QO21frXTfJVwWgDeivKIVK2Eui1QcTgEOkHtOqjbQVndzo+vX` +
		`3L6lbJzy+v4LFodbQ7AV66mbZfNHF1bR8AKsnzKENX68UZXPk+w/WEcFwpmxgETrqV28` +
		`il8iPHv1DrnTUIHd6RXEDoHLfvB1lhBtk4ewroRczgoYHFQ2LBu8gmR+hGxAL55AR2qq` +
		`zjaBKjrfTwf3mXi39wz4n0GB0KEwvv4ePIZ7Kp/sy0svam6+sd2WAmEsngPVqw+KQAaM` +
		`9zt355uNvCiaQufstsGhEVXDqYTKtbWQOtYTDV7qaQ/FbS3Pwy72EQXetv/ygYXcKvsI` +
		`/568Q2VHG2Dt4Z9LKUj59ovotR6qb1y9CT7QOvl1XsjwEgb1AuvqxD08suqSARgVc2QT` +
		`0aAUyBVfTSnfq+1Cci6FylS1jZCcoZpPIFRwFcbVdTQOjWvCrhLk3P760Ogesxao9ajD` +
		`54+APGcWG5bza+x1F4AFT/0TknTjAaxPgRaj9astWudml+80540V3+4K9CQEVD9KdRug` +
		`vI+UKN9hCHYAcK1DNr02DdO2fTQJzV17H/qZBa5QEgEYPsl3k1/yiirbsfKp4+xv0ye6` +
		`dp59KioJUExbM/DNGq4x2LVH0+yH5BnutIi2Gkmx5iAzdwafcTqqjFlZXSijv01hn0Vh` +
		`gP3W6x99iT70vkmANbvRs9ggBWkTXUtK58+tSkvFBprW8aRstknDbYvk29ApoUl9edW/` +
		`XILq2yQSbLJrvVLcnCt4UPbXu4Tc4/RWhYip1cN+22J2iOqYIuk6mrm2Wq7PhyiraxmF` +
		`xvtPtIEJHFdWs/HVLHL9oeHKLfX+j2C9kVMkcJwklZ6cSNjJadEKnQ8AOull36pKLMkb` +
		`Z+Y4d2xBAwBxyXXzhZopaPGViXSB1rPWGBSv2jvvqS3JKvmjgdgF+DyRLv9BH59ZARKV` +
		`eKIPoL0oQNuYSD6QyoQ6zYI7e0QqDgwEpdQ15rVULMRTCgigStUSywiB0LhHQ9eO6fLc` +
		`0GDCVgsfeIk+4sY1wst6dnhCI4OWrSlnDUTj7O/Zr5o+EzbcSUGR9Neq4oNgRAHCQzxB` +
		`qyrYV8gzPpwGyrt+aOoupZdT59qFwjzTKPeo49CV3WGfbuXsuSFqyLXaEgqH96Wo6wA5` +
		`dISaiv4bMpAuyZxXACPmk9PquljguyacHzTNYyczaHGont4H+smnxGR5NESWG+cfeJoI` +
		`0vZxw460ZOuTV5Mb/tayyJbIvZmoG0bi1xga9xq9tmS1gWq/q45NXe05HYBLJBqrGhtQ` +
		`Gr6ZvRSS05G1HG9TpETR+4l9nrZBeAE14402o4fAEtiy9/nhppi1XQ9Q+XCaH4VxnN9L` +
		`Pqo6OMigGldsf30YRVwpLnE3uAgqSjQJQKwNiPQx77TksJ6JwVwzal/NYaromue2JvT3` +
		`A+J/nZ4pDinigFUr4ZarT8E7RRia9RvycjZ5mIDnWvr+GzKKZGXtYkEtjCj53GkVUdHa` +
		`tgwcahd/MsXyfMcNk1vudu7l+XPnck2d854g5aVsfyxQWwbuZBuZn9jFVFhnd3LmeeC0` +
		`Muihs/kOHnA95ax/LlBja7Q4Np5tK+ojXzBkwHYBnxHBukD+NkGFrx6YVQOkoUZMYdBV` +
		`pjg+o0sdP9mvz0Wlt77CBLssuXyPu+2bdBoiedlM24b15btZpdYj1xG7QKuO31t9Vdqb` +
		`7Qd7PY5wL5FKVKFyQVwO7pxYIw6Llt4Ax9jCHAIQ+xIaGOFXQDSS6/7hbwklfkBsN45C` +
		`Q2tM9EdoFKFpfrKo+36EVwAy8aMvaVBoJcbQiRwxCtW544ZTwIn0jYEVllHsR+T6Ceqv` +
		`4lX7vUpUtMF4qjYt0pyfMkBcKgrlDc9UDxg7R1nnLb+3v3VQdY/fXykgHUsgEfN43Bq6` +
		`WYZNk86IbUTnsPn0CUUblpvz533xpkcEjYJnDchNkwczJob5nNwsDrGhpMgqeGTPwz1d` +
		`pedQHdEp4iKt3Qn8/59btPUg5EfMFDqdTIAa4zrZ9KvzNAl2l7W72+cSeew4YhY6ewCW` +
		`EfD6qt3uxK4Iz3olqKJsY+dbOQT2nMA3Z0X3HVI9WGwLZ0Fvmg72LV/pXa74SoXwApjd` +
		`Y2xZfWR0cdGEl7SWeljKqfeQC6AVcIrrjYV803yA2CVw5M6Kxs5UXFOt020He4CWCprr` +
		`Pkm+1u6qkg1FOKRO2Y8AMeq69H9XXVZJpF7Fak+GFLXtY9417eofxSAtdyFqyG4EUI9I` +
		`3ZwDHXZ9eb3r/jkvKnRkjUawKEq9rnq5pMn1G82wV4T//q6d+hdXlGfydKooZxRcvQ0N` +
		`2hzbaJt1WQLu/Id2nau4FgrSHjS8REfUiyN+JC+gRq6egHwlTPp3NlwRK1F9ZSTGq7fu` +
		`PZ9Dq8I0S3WPnYB3Ch1MuJo2kQFHam0nTD+SV7h1cy1bdc+Tq3oNcy37dteHGt7jfV32` +
		`dlKAnHtXzmo9HeXXADH0wYaANyG3gywi802PvObTQAvdVTQaCDF8sn9IESD1QVwPOeap` +
		`LUkrB5L/AvfIhlXibzQclwlel66ikUgldSX9Jek1Vz6aEjSyz6WOejazvKqyzsdk312+` +
		`2r4qQBcCeVNnYMV+xa+eO3igb+MB2B5WKUCiktSVatrWepV0sVy1pXizbyZNXWG2lgbV` +
		`2vAsNdW3xOQH83AVZFr66ibckq9+tRo5OHvclionAO8AFgdXZW8XhOQyTE3IuVjPdQug` +
		`OvtXzVyVV2pppKGXtIkZd+6UtNd/QY+YT876c1xNmDXsoBoqS7pKWeW7OC9bGcTy5rY3` +
		`V4BrPizFNnGt7JlE8DpSuB4APYigWV3y/6OJ4GVPZboJIIrgeXgUsql1H+ZATItZKPHp` +
		`sK6TsAYAFeshlkb4cj4AA5UrVgzYmHfy+IBWO+FBTtMDbts/ybUBCwWp5IBNeIDjg4E4` +
		`t4EZb9+VpiqnW35uPNuelhButRZhMos9poga+oCHChnUCIA+0n5vGQm7XoajrHChCedn` +
		`EACf0DfQMCbBNaaRsyml8JOtQG2TBnCStdciPdhMRaW3iTbte2S4r+r+RCLOk92cJgw6` +
		`1hkA13S1HVaKRFEdq1sWgF4K6s4kH52vFi0jc/sG3plB8vrLW+zkkaiUzK9AjjIgDh12` +
		`rIJYJl38tE4Tp64b4M8uQrpyDfovi/NSWDZoO6dRbrlIV4+vlRsAT0egCVJIzdjNKZoG` +
		`1jajSRtsg+QG+6Kzf9eDV/fCE/FB7CcXF/4dMwFh255dZHrnKp3KFWzd+JptnuberUyx` +
		`fzoG2ZzbKx0jbPxGsrZULWW3cGjqGM7bcoMPZqTvhrDT1ZWtDNtxzLmxnN+XT+DAWVB2` +
		`niVwAqvddjPwFCQ8ITjmTdiPodLBd9RwZIXBtrhhHoSgONW3ZCDSWmSUucEShd0TV+N6` +
		`EysiPrreorlaRawO9DdBrAke8S+jZgnrp2sZBEBOCKNj2/kafYK4E8YEKfqRjYBLA+u1` +
		`GjtRWCL9SYr9iq1VeCJdqI1B2BxxZWU8bzQ8k7b4b8EANbv49n7rhc62vMdz5Z2n67Wr` +
		`nmkZscekFgNoY3wcSUcHie+Xr2aI7dMHnP2p3e9EgvgaMmm2Gn7WRwr7220xzrudzDOL` +
		`13HkNf2ftopzPTkUO/2+YjZDA5YhHZYLH3BCZ258904k3ZhY6s4Sb3Q0WscPocjQ2E6K` +
		`S7dLsTB1FAz6eT6h98IwEpC132TjSiSC604ayTPWTFYJV240tXNhRbQq+186RA9ObJJy` +
		`qUbD1a2lmv/uhNJ0n/KXNrRmb3ssPOoY+POXgH8dwZwU5OnlE0AazLFgZVwERsHFmgFE` +
		`Km6cSSYHUZKlGDiAjw2DiwwSX3W89B/x5PA+ojItJOjyv2gaC2KAwvE0R5zN0wUq3ZLC` +
		`9A8Cg/JDpa0bvBJRGLrG+GWSvhtfAD32vnq3RcuGzOxOQCLe8o+2rc/YrTXlrHkmcG2y` +
		`94Tud5aT439NgqxdeLghqhMsu718dkQuycPssMiduK8Gwc2hjI51bxKYPX95od0bVdDX` +
		`9nWiiEncqxJAtcncMQuVKqxMqGUyuiSgCog1toOkUiCvxxPUo31byxtYoVzmslqZP+67` +
		`SKeZ6mE8TOzvAL4VwxwEjiiV5BtAEdnYulFlxfYzZqStBQ/9H5Gq8LJJLAAJC1T/JV0F` +
		`8A0puaSiqw55FyKBrCbaCLTRP3VV+3UXyDUWrQGCQJ3LZKy+gDJE60YteZRX2kW+gjI5` +
		`SwprGcsB5f746y/bSV81jvy62iqXk2XvTOfumLJ5eOSAVjdXHvPtllXsNhr/FVqZqc6j` +
		`jK1nuKDyTBY/3etI2j8nUzSWjpWc4yksLzHoTp214axAhV0rguzW4UEZHP7ATCK+85gk` +
		`IDvxKYXvHBV01CBAKzoXsOdoHG2qhNJUn8VdxVwJTmlBuuAg8JE7Zs8xYZB3LRK/Sba/` +
		`nVbuLFf/X900of7d68A/hYD7Mv4GlO2AazZBATZwtG50G4usbzPsc6hZADWmAKgJKTGd` +
		`M8lu7nQcj7JtIkOXwnQct9JwirmrL6S/m4utNRgAT42riyBp7au4NO6NYablaW5tDd9Q` +
		`BQqk6YRtf6f925a0bJ6NW2rlky9etEXrvECYO22PqPKp+Szs7kOo1fA0N0EmmSieAatG` +
		`so7XWXY1Hcw6xLlITc34PA3aRPoQC+CdDR1BMJBqoLVbJk4lI0jZ3OIDeCo5BYvhxlGz` +
		`aMPtfRszqkmALt3Y/racKE1HuoUr87OugrpPHB2dpjSqMK58Nw03+R9jNEjyQnJng7No` +
		`mu7arpYbWgvSehlYsWC64LsturYvm8V271Kfy9jR7cZ4Xih6wyfPnmCrT55IjdUVhtgW` +
		`aJDHgKwgoHeqyB5mjr3jZSrlr3KkyUAJ3yiCrxf1+SvSzAmOgE8py+ErdJup8J0pSxcR` +
		`TBYHZHOdeWEA/upq+hE9SaofmFgwuwVX+uV9A32oN1Oyx6zkZdYqvCo2XZooKy8giW/j` +
		`/17gpk0punIsRbURCd0xDYXgKUjeTvx6WtbuW0sl5DvbHjPS3TVxubSJj0P1rIaujnSj` +
		`Xe1CmMS5Ze2rP0D106nR0UbDo5X2cM96BCbWdUcE1znV7CMtlYbVkwaYB8/i0sCsOsqL` +
		`WrGykK0j75klCR5xR7ZonLmxAvtZHTC4htMzm6Z5I1FwAaMcQt2Fd+e/K446nCGEjoIs` +
		`Ev/2gCso43+e38dy5Jlnak4gRWgvQnTUZ7ncBnbJg9q/gprAVhxEC+Z/n63ldP28vdmv` +
		`nysK3llUikRI9m54pxuuXAmexD4XqPlVGGMl+TzwtlDuiu5i8C1X6F7xV4OCFdQrpRRq` +
		`b8E2bVnHxteGJa8wpNrKyvmWx1my9OnsTZZcYESgO0H56YwyoqWS0D/SqxI+ip+qh83l` +
		`CY1WmEW/Ujuyxsscy/6p3GVkXTfjYLvr3MXM1s5gPP0kFqpDayY6ttO1TZVbtMRPvcYX` +
		`bpPQv4SFTf5vPNzerr11tJdUG76K1ekofZD61Khc8PhhDZwo6B3j6gAABM4SURBVIMMe` +
		`VxLWlM3HOSPN4xiqTMcwOpUuiJniY7PpbWMOJ0VY9GtxAK0PD66u6SwamhnZMfjgHvrR` +
		`2pVTqyM8C/FQVp4GEkpgX9xfnTFcb4L8Mvjo/DoV1VUJU7p2hSfYiF0Uzq4ew5DNkgew` +
		`0iFwI5craGFJnLokMF9wLNONlKu2OlnHqU0XgP8qCWE4SMbl5sgUgUop4kcfrje0tp6S` +
		`qUshk1fBrxsH8fUjdXKDs2UTZvt3ctm1upVyfqkbE+W3fH/G7jKnuJljNGmSpRlDgjAq` +
		`oMuMVDEtJ4fcysP8HwR78HApcNhxn2wvkhDqA3n2h7AGF0SUKIsc0AAblQPK8vzZXh4S` +
		`dlHgDuYws6YelgZnioXw6kq2Vc6we0/h99/K06lyFwsIo05dPAqks96A8aoTFOJssyBh` +
		`Af6szxvBoZXksU3gWn2WP+GOAf6MzBNLod41/Fxac73zoCrnoPVGa1bnv3dKHxeyWkYI` +
		`1d/ibLMgbgldbI8ZwaGl407wqmmGBlOCd1KqShqUnHI6Ht5d3SF4ZPhpSIyJ5+xfXNdM` +
		`cbzFSlF/czyvHi3qJ2SfRNW6cvzGqOm10GPHyS8OLT5WHDh7CLuSlQ2SxH5eDRhLNz4k` +
		`MfSuPndZ/cR7No82egAcolywAEXwBHFp6BJH/QLnISM+AttWla2oDfUeHHXg13YOhHNG` +
		`gqffxW2R4vowtvfl49j3cvzTaJC2YW34CJfkQtgHVeKLY5SQFuTvXuRUyQu8bKaFnYvo` +
		`C0kW4ocWPoCNUeLBsG5r8HGxhfCJRs6l38f34FtY3eb2DthcrmEVjWXbGCdRNKJpAIlJ` +
		`WWc55R+aH6JTa9WKdAtxVuWLrbxcoJ2xZHwuX+A/i1AcrbRxmAyfzisAPeb7yUJwKr47` +
		`l7Ik+/1xMyvsqzKIVYVR2+kYwQNGX3e+uS9la5N8r5FWNcDznsNPlL+YuFQ1It0vMEU3` +
		`WMoHE56X4kArNwZ5dAUGKks65lO0TXvS1NWinKaio4E4ObvUW+8pQ294LS3YEXsBWn52` +
		`/n9gJO9cZXBvJC/lbSemQXgnwH3FNaWVSlR4BWI/ZFOtanAV76PLfhbNfCEbnj22WtZP` +
		`zj1fdgSuckxnyRnii4wdU4U/sxglNNaoixzQACORO4KhnR4XsfvVLM4NcpuhcrU1pS01` +
		`8XA35O2atpgzilw5r9gT36jgKc6hzSdBT5jME1L3aWwvVKX5jkgAL/mGJoFwqtLU3yTG` +
		`5ZflGq0jg3r2xXvSuBkT+blL8NlSm7JHylKfVvD9K8ZzPn5W03rmVkAfscReQWwa10S8` +
		`d2011G0arQq0+v7lQp97yF4KApCqYyRYp84de2mG8wZKQ5X6uaDAwKwvIUF4M6c5XxHd` +
		`JtA+qSDMY+nP0xuR2hc1cLf3LVlMGw6fHCyv34ZaD0GGN94nLkGMyQDQ5eGSMIBAViB1` +
		`jy7MlXRWVaUrrTMDMmhonv6VEynaOgGYHIaq/3sMBg4J6fZWqrXqQtPY6LSHxtM0V8Wk` +
		`MaTyFlXAbgA0iizkwT59cKMjyV+uLIa/5nms//D9XBN7k7yfQ34U9MlrzGYIjtGlSbf8` +
		`9RdAM7zQYbsHQRUpVOVkCuakJJUBvf633ReiDPegOnnpjOCp76yfcXjE5q23mow+Y9te` +
		`dpFcTcSgPP4fivbTvk72bvVqKiksM4p6EhhuvTR0TBoHtRkt15/M+cvqgytrLB7us8sx` +
		`f4CsFCUSvAixSmju90F3J2BcRIPoWMQEmyK0BQ8KZSrGw4ywZIH7ocfZrWqjc52HoYxD` +
		`ZdHOwy2sH5iMErMKlGWOSAAN9QSzPJkjYeX1JX0zX7Oe2aCUzlgjtw+uhgiE9+1dh3ho` +
		`4VwcNZM0e9ijFjbiCwsrf7bBqO7aEqUZQ4IwHkq7J4dx1U8fukOeuV2fZBlZqY9/NlOb` +
		`SANlAkQj7geJmXFoaWY36kYI9bWkwPeO1WawGB0E1yJsswBAVhuE2mZOSSlWhyb0yCPE` +
		`jPl0NKlKgVL3wCei1pduiAOhWDVQuilanMZI7HwZIxR9CgeePW7+QZTALkFGdtzwQ4kA` +
		`EeXUsvRQvOTZlHwB/5l/+qmv2hKF8Q3j4FHY9Is0nvKTSpORkled+R3DEaXQpUoyxwQg` +
		`BV5zGHeqsLOyhvRjYC5p4L2SiugqsBqLKUF4rawbilUZqTKzVMY06huSBzwavX/azBfy` +
		`v3TbX0zCsC6POgrudt6fo8aqLKWcvwS1Y/LHR9iZlIcQItLdDNpOiD+z4fgP9LOk5bKL` +
		`NW53gpJAF5t7EWD0eVPJcoyBwTg3wM3ZXmeqOHzXzNjsXNPoHSBgiHJq1eSrCZVEA8eD` +
		`HPTKrqiqr2q9Vz/3WsGvNrEIwbznYLhbQteiAAcz/LK0pb1EhVGjnvBgdjrgf5UQbxgD` +
		`gzQx9M3CbznYowCXDYlAa+a3Gowv/U9U6mDbw4IwJcAf/PdM6UOUuN+k1LPbHQSiHXmT` +
		`SU58046z+G12GQqIP7xj+B+mS++KBXwaoILDeZVXzOVGqfEAQH4GOdASUoD+OtUeKXXF` +
		`dCUBy+vINYR6F/746TvOPExR8OieiHqZTKx5HSfktcd92iD0YGwEmWZAwKwrovPwam7w` +
		`i36qnRLFa7VucqcU3vg0/p7df1N71cSr1sBlV5q16Kb0c/BmDXugjyozW5TJXdUGBone` +
		`fjbWKm1Vw64hd11EPcQr51Sa1fYUVgVhZQtMSO1zaXe615AB/lTJT8gfvZJuDpZ9Xj7g` +
		`NGFGLMpBfCqy3KDyfP58lSZWXz9XAC/7gihLO4gd6mTqW5ijwPiN1MdwG8/2bz6dLbx2` +
		`zGmvVcQJ0+tlN16hcdQUaJF/5/B6A6cEuWAAy6AHwOUHpVFKjz7N95mVdBH9VB/ATRK9` +
		`M0GZwSXL2ZoYC8gPvRQWBk3Aq6t3gfcizH1NY18qM3Rm/idwdySoV2VhknCARfAWb7ku` +
		`/gu/9Stw7p9WPZxVkhwyXQFei8g3rQOuje6x+4z4CqMUXHDekoRvOp/vcE8nRWelQZtw` +
		`gEXwBKPutYmS5S9qhtZWrA9rEqVyDx9FAhnciJd9/vXTA4YNVYyEE97E04/Rx20JWlet` +
		`2OMtpoJ8GqMngbj56KYLDGidQxrA1hkYckXmqXDo0qmz2GyV4af3Wzn0gT9mzYd51RAb` +
		`5f2SIkHaA7ETz4Jw4dHtmRMky2lIXm1nmUG41zOkMX9lYau50A0gOUmVpWUzNOpX1/KR` +
		`/99BDsyP3SuRpS4Uplaab0pb0OunT8DCh1lm+KAWLdu/2DkyJV3TJzYD2OaKBVpglc7m` +
		`mwwI7O9tdL4DRyIBnD27OAnD3uPi1cNRUVWpLTl5yBSRp77ZkhtGz8ElAhVz/GMLKf5Q` +
		`RwQ63yEajf/FOh+5pnrzNtvNzmalAHwai3XGYyu6ilRjjgQDeDs2cHTOs3n9F1SHkEnC` +
		`P6r+IHseRuqK6dD+pfn6IlGTWMD9y744d1wkPv7447bbebPb3SRUobAqxlK9m+OH3Mje` +
		`WBhySOZkYOjjfaxoM1yBlT1bfQ7IUDxmolAdY53ncHpEm5DsV0pk6or1yeDE3oYypW4E` +
		`vo2cKPV6aP6VpmPl9dHnjMI3pL96+HZZLpJLIB1giTzMbzVoU30qUt8++1bwJM6RQqo1` +
		`mEuqcyZLAM3utjb6AAvjoXdP5A8yt1GZFZfAcgOsn3MseSCuE+fWrNmjb3rDIJXw/3aY` +
		`L6fux2XZhIHYgGsi3Uyn024PbCDzlbnpCyXJNZNTVqB+6PilZmqXK3dqvKjimO5PyqMq` +
		`Xl1FlcOJv3r9yMihfRCpyzCRVDdIT/bSFobWCB+pEfYbN0YzDB49WiHGExah46Tvh+lB` +
		`k040MSlYmHpyEpm77WxjGrHplZlXA4vrSj6R3qrqleo2LN+3MrEugtAYNKPiqRLf9RO+` +
		`jv/qo5eoooX0azR4crpgFL5NZcMC/2IWzIwpBJrbP2re8gvSv5m5WMbcVd1f6iW22uUa` +
		`KbqkZmiBQYT8XGUKKcciAdgOSsze7t6VaCKcis1AOeUHa1gsmpTRUU408/iRwbzn62Ae` +
		`wW3xXgAllxRUkfmAh6bghvpHvZ6XL3gmNSiFrQpsJEedZl8FjJw5H2uP73UovhV4JuJC` +
		`1ILaypwVsbWvqxsDX1rc+yLzdjqW9ZAy0JrOKImk8/idYP5fMtiUvHsJhGA072ptjEH5` +
		`rZdweD9pUr9hfBezGmzghP2ZfJZlJI38vhcEwFYDk25cBKHfvws+s0uCzhnx0A/XUpts` +
		`8SB/+uykAu2ZeomDiWm9TaYIo7kZ4nPORo2oZ1rYf3ISf5LfylPHP4+o1eemv5ApRHS5` +
		`sD4w2dw03IF0TJBPzSYX2VioNIYqXGgOQArGKPgSaO0u5SmGXnRm0x8Jfs3Tqe0uFbW6` +
		`YaLpzLl77pGLV1SIO9gg/EbNU933lL/KA4062m2sBQvvD1tjp1wz1Rm3ZmJlybtpbT6A` +
		`Y6/exrz7lA13XTpHoPJZCw53fW0yv7JACwbWFWbvKQ/JGZg97/NY9OXlfNUonxzoOu/P` +
		`mT759K9OVB5Kb0MRlK4RHnkQNJYr4WVfn50cO1GantnMvaYR5YV+dTBrVsId1XOWjr0k` +
		`MHoRo8S5ZkDXgAsKaw6wenZwmtCa+hdl8n4Y55ZV4TTLw+toV/aMWAlr/Y1GHmgS5RnD` +
		`iQFsNZnYelsTXrexvsGv81PPlTmcInyxYF7Bk/nzrnp3tt7m8EUzv04+eJlgczrFcA6f` +
		`qbrJVOvdzT4nmnMvTMTzpMCYV0RLmPQz99h/k+HpbFyHSkZWLp1IQ0OZrirJwA7UvgLw` +
		`P+lPH9w2SZqj8hMYkjKi2jlHUNLN1PXr3saXDjDYHROq0QFwgHPAHZA/DJwacprn9FhC` +
		`SfvyexRxZQX08o6Tu/wMWfsOiqNXf/ZYFQzoEQFxAG/ANb9Sbp1LrXjaJde/yYvP11K6` +
		`MjHC3DRDVN5ZXKqsXid5z7SYLJW5z4fLGkJc/oCsCOF7wF+ltLmKxatYv8AfQR8z5vSf` +
		`KVOLgcsKhavpbp/7xRZcofB6Bq2EhUYB3wDycIKOfUqUsunja5QWWDMaLHLea3LQs5P+` +
		`QDDB8Awg8nBFbQt9glkbWO+AexIYcVz5ZX2Hxu+cMRU/vFkqqpc1hjRogc+79tTefORV` +
		`HiuTCt5nZUTX6IC5EBKAHZAnNoNP6GPNlBzjLKyUp67APlYyEuyCK3dRN1BqWTCfcVgX` +
		`irkzbX2taUFIgtL937d7JuJ406dyr0zUpEIvqdq9R1+cupb/PK9VKqrPGIw32n1/CtwB` +
		`qQLYB381wVZ/g6Ily9aw76BlQRse7pE2eJA2NTSduEGqo/x67xScd+TDQ13BWdriaVx0` +
		`+NAWgB2VOkjAdUD9nff3m+OfofvfJxOVlB6O28NvR/s/y4/WHyaz60q13mQwcS9CdznW` +
		`KXmWeZA2gB2QHw+8A/Au0RtN20Ze87UdSsZWUOW+VSMw1u0n76cvcP6+Vi8SuNcYDC60` +
		`LlERcCBjIHHwvoq8DwQ8Lzvh/u/w61LSlLYM8N8NLx/0NvcPs/P4ZE64DKD+R8fs5Sa5` +
		`pkDGQOwI4n9XVEaXLmRrX3b0snyH47KM+MKevpdZjddNlYT7t7N4zpV2/lqg/mjx/alZ` +
		`gXCgYwC2AHxrcDDnvd37i1TeeN3JY+0Z4Z5aHjWd6bx9m/8nPy6yWDGexi51KTAOJBxA` +
		`DsgVtrdOE97NXV1vNdlMafsLpWd9cSwJI3e7riYs7f1xwp6fbZ3GozSY0tUhBzw+pB9b` +
		`83CUjE8FcVLTm1nL2XrSb1oY/nzZCcfuXW12G/20mX2ZqqOV765FxpnMN6ekZfRSm1yz` +
		`oGsAdiRxNcDkz05tk66+998cFfcq21zzpVinfCEe95mzs+8OK7CgG5U+EOxbrW07ggHs` +
		`gpgB8QXO1d3Jz+C+PujpnHTJ35st9JzdDnw62On8v2FXnwJ+4HLDebVEvOKnwNZB7ADY` +
		`tVhUnhCt/YmJrNtFx/32MERpYvQfL1aiyqWMnBrL6x2yUyQHcDnDWamr/FLjQuWAzkBs` +
		`ANiXbP9hnM1dmKGVMxbxbIT29C7rmfBcq2QFrY2uJFDP6qm9ohkFT/XA+cYzJJCWn5pL` +
		`elxIGcAdkAsUCrW2HxVDoF41Ylt6FkCcbOPd0NgI4fO2UfVoEOTvAZvAt8wmA3pvS6l3` +
		`oXGgZwC2AGx5lSZWnk/Ve0yPrWZu4w1J3bjgHDXQmNaQaxH4D1szm72D1I6aiKqccJ5v` +
		`zIYJWuUqIVxIOcAdvlnYZ3kOLcSS48O0z7hk7O7U1kCcaP3bn1gG4fP2cz+QTpIkoiWA` +
		`18r2bstDLEx28kbgB1p3AGYBFyVkM3li1ezeFCAvrV+j8S1zCf3SflKBs4LUd2/OZv3G` +
		`WCswextmUwo7crlQF4BHCWNv+6kX8Z3XIVWbGDmMfsZXJXM1mvZT3ZW2485dXln6ioTO` +
		`fjkqLrFYF5s2Ywo7a6gABwljZWCeVvcxxPYvIM/H7uUyzad2Cof3wuVs/naR0didU508` +
		`OO/AKVFlqRuK3pBCkICR/PbwlL5l+eApiqz8qZ/Nuwt7p7RumpLf3/YNB566zSsYDDOu` +
		`6mCc1cZzDut6L0tbdXhwP8D5MK/fajAwncAAAAASUVORK5CYII="},"duration":31}` +
		`,"touchSupport":{"value":{"maxTouchPoints":6,"touchEvent":true,"touc` +
		`hStart":true},"duration":1},"fonts":{"value":["Calibri"],"duration":` +
		`132},"audio":{"value":124.0807470110085,"duration":30},"pluginsSuppo` +
		`rt":{"value":true,"duration":0},"productSub":{"value":"20030107","du` +
		`ration":0},"emptyEvalLength":{"value":33,"duration":0},"errorFF":{"v` +
		`alue":false,"duration":0},"vendor":{"value":"Google Inc.","duration"` +
		`:1},"chrome":{"value":true,"duration":0},"cookiesEnabled":{"value":t` +
		`rue,"duration":12},"visitorID":"f385cff34b7bed17244b93e12ef6ee87"}`

	d := json.NewDecoder(strings.NewReader(f))
	d.DisallowUnknownFields()

	var decoded Fingerprint

	if err := d.Decode(&decoded); err != nil {
		t.Fatalf("Decode: %v", err)
	}

	t.Run("OsCPU unmarshalling", func(t *testing.T) {
		if decoded.OsCPU.Value != nil ||
			decoded.OsCPU.String() != Undefined {
			t.Errorf("got %v, want %v", decoded.OsCPU, Undefined)
		}
	})

	t.Run("Languages unmarshalling", func(t *testing.T) {
		w := "en-US en-US"
		if decoded.Languages.Value == nil ||
			decoded.Languages.String() != w {
			t.Errorf("got %v, want %v", decoded.Languages, w)
		}
	})

	t.Run("DeviceMemory unmarshalling", func(t *testing.T) {
		w := 8
		if decoded.DeviceMemory.Value == nil ||
			*decoded.DeviceMemory.Value != w {
			t.Errorf("got %v, want %v", decoded.DeviceMemory, w)
		}
	})

	t.Run("AvailableScreenResolution unmarshalling", func(t *testing.T) {
		w := [2]int{812, 375}
		if decoded.AvailableScreenResolution.Value == nil ||
			*decoded.AvailableScreenResolution.Value != w {
			t.Errorf("got %v, want %v ", decoded.AvailableScreenResolution, w)
		}
	})

	t.Run("Timezone unmarshalling", func(t *testing.T) {
		w := "America/Los_Angeles"
		if decoded.Timezone.Value == nil ||
			decoded.Timezone.String() != w {
			t.Errorf("got %v, want %v", decoded.Timezone, w)
		}
	})

	t.Run("IndexedDB unmarshalling", func(t *testing.T) {
		if decoded.IndexedDB.Value == nil ||
			*decoded.IndexedDB.Value == false {
			t.Errorf("got %v, want %v", decoded.IndexedDB, true)
		}
	})

	t.Run("CPUClass unmarshalling", func(t *testing.T) {
		if decoded.CPUClass.Value != nil ||
			decoded.CPUClass.String() != Undefined {
			t.Errorf("got %v, want %v", decoded.CPUClass, Undefined)
		}
	})

	t.Run("Plugins unmarshalling", func(t *testing.T) {
		if decoded.Plugins.Value == nil ||
			decoded.Plugins.String() != "" {
			t.Errorf("got %v, want %v", decoded.Plugins, "")
		}
	})

	t.Run("Fonts unmarshalling", func(t *testing.T) {
		w := "Calibri"
		if decoded.Fonts.Value == nil ||
			decoded.Fonts.String() != w {
			t.Errorf("got %v, want %v", decoded.Plugins, w)
		}
	})

	// spew.Dump(decoded)
}

const (
	testError = "test error"
	testValue = "test value"
)

func ptrToString(s string) *string {
	return &s
}

func Test_stringValueOrError(t *testing.T) {
	type args struct {
		err *string
		val *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no error, no value",
			args: args{
				err: nil,
				val: nil,
			},
			want: Undefined,
		},
		{
			name: "has error, no value",
			args: args{
				err: ptrToString(testError),
				val: nil,
			},
			want: testError,
		},
		{
			name: "no error, has value",
			args: args{
				err: nil,
				val: ptrToString(testValue),
			},
			want: testValue,
		},
		{
			name: "has error, has value",
			args: args{
				err: ptrToString(testError),
				val: ptrToString(testValue),
			},
			want: testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringValueOrError(tt.args.err, tt.args.val); got != tt.want {
				t.Errorf("stringValueOrError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_standardStringValueOrError(t *testing.T) {
	type args struct {
		err *string
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no error, no value",
			args: args{
				err: nil,
				val: "",
			},
			want: "",
		},
		{
			name: "has error, no value",
			args: args{
				err: ptrToString(testError),
				val: "",
			},
			want: testError,
		},
		{
			name: "no error, has value",
			args: args{
				err: nil,
				val: testValue,
			},
			want: testValue,
		},
		{
			name: "has error, has value",
			args: args{
				err: ptrToString(testError),
				val: testValue,
			},
			want: testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := standardStringValueOrError(tt.args.err, tt.args.val); got != tt.want {
				t.Errorf("standardStringValueOrError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hardwareRelatedCompatibility(t *testing.T) {
	deviceMemory0, deviceMemory1 := 0, 1

	type args struct {
		a Fingerprint
		b Fingerprint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic, empty",
			want: 1,
		},
		{
			name: "special case (a has errors)",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 0,
		},
		{
			name: "special case (b has errors)",
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 0,
		},
		{
			name: "special case (a and b have errors)",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 1,
		},
		{
			name: "only a has ColorDepth",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "only b has ColorDepth",
			args: args{
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "ColorDepth",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "ColorDepth, only a has DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 19. / 21,
		},
		{
			name: "ColorDepth, only b has DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
			},
			want: 19. / 21,
		},
		{
			name: "ColorDepth, DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory1,
					},
				},
			},
			want: 19. / 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hardwareRelatedCompatibility(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hardwareRelatedCompatibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a > b",
			args: args{
				a: math.MaxInt64,
				b: math.MinInt64,
			},
			want: math.MaxInt64,
		},
		{
			name: "a < b",
			args: args{
				a: math.MinInt64,
				b: math.MaxInt64,
			},
			want: math.MaxInt64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarity(t *testing.T) {
	type args struct {
		a Fingerprint
		b Fingerprint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic, empty",
			want: 1,
		},
		{
			name: `special case (a.OsCPU ← "", b.OsCPU ← Undefined)`,
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString(""),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: `special case (a.OsCPU ← Undefined, b.OsCPU ← "")`,
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString(""),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "only a has OsCPU",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Linux x86_64"),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "only b has OsCPU",
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Linux x86_64"),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "OsCPU",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
				},
			},
			want: 125. / 126,
		},
		{
			name: "OsCPU, only a has Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
				},
			},
			want: 37. / 42,
		},
		{
			name: "OsCPU, only b has Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
			},
			want: 37. / 42,
		},
		{
			name: "OsCPU, Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
				},
			},
			want: 1541. / 1638,
		},
		{
			name: "OsCPU, Timezone, only a has CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
				},
			},
			want: 151. / 182,
		},
		{
			name: "OsCPU, Timezone, only b has CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
			},
			want: 151. / 182,
		},
		{
			name: "OsCPU, Timezone, CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("Alpha"),
					},
				},
			},
			want: 151. / 182,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarity(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("similarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExperimentalInMemory_Do(t *testing.T) {
	var (
		fZRMtpX = Fingerprint{
			VisitorID: "fZRMtpX",
			OsCPU: OsCPU{
				Value: ptrToString("Windows NT 6.2"),
			},
		}
		screenResolution = [2]int{2560, 1080}
		truth            = true
	)
	tests := []struct {
		name         string
		fingerprints map[string][]Fingerprint
		f            Fingerprint
		want         Fingerprint
		want1        bool
	}{
		{
			name:  "fingerprints and f are empty",
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: "fingerprints is empty",
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Linux x86_64"),
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: `f is similar to fingerprints["fZRMtpX"]`,
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Windows NT 6.3"),
				},
			},
			want:  fZRMtpX,
			want1: true,
		},
		{
			name: "f isn't similar to any other fingerprint",
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Linux x86_64"),
				},
				Languages: Languages{
					Value: [][]string{
						{
							"en-US",
						},
						{
							"en-US",
							"en",
						},
					},
				},
				Timezone: Timezone{
					Value: ptrToString("Europe/Warsaw"),
				},
				CPUClass: CPUClass{
					Value: nil,
				},
				Platform: Platform{
					Value: "Linux x86_64",
				},
				Plugins: Plugins{
					Value: []PluginsValue{},
				},
				Fonts: Fonts{
					Value: []string{
						"Batang",
						"Bitstream Vera Sans Mono",
						"MS Mincho",
						"MS UI Gothic",
						"Meiryo UI",
						"PMingLiU",
					},
				},
				ProductSub: ProductSub{
					Value: "20100101",
				},
				Vendor: Vendor{
					Value: "",
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: "f isn't similar to any other fingerprint (incompatible HW)",
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				ColorDepth: ColorDepth{
					Value: 24,
				},
				DeviceMemory: DeviceMemory{
					Value: nil,
				},
				ScreenResolution: ScreenResolution{
					Value: screenResolution,
				},
				AvailableScreenResolution: AvailableScreenResolution{
					Value: &screenResolution,
				},
				HardwareConcurrency: HardwareConcurrency{
					Value: 8,
				},
				TimezoneOffset: TimezoneOffset{
					Value: -60,
				},
				SessionStorage: SessionStorage{
					Value: true,
				},
				LocalStorage: LocalStorage{
					Value: true,
				},
				IndexedDB: IndexedDB{
					Value: &truth,
				},
				OpenDatabase: OpenDatabase{
					Value: false,
				},
				Canvas: Canvas{
					Value: CanvasValue{
						Winding: true,
						Data:    "a347f8567cbc22e94f1424935cba5c3f",
					},
				},
				TouchSupport: TouchSupport{
					Value: TouchSupportValue{
						MaxTouchPoints: 0,
						TouchEvent:     false,
						TouchStart:     false,
					},
				},
				Audio: Audio{
					Value: 35.73833402246237,
				},
				PluginsSupport: PluginsSupport{
					Value: true,
				},
				EmptyEvalLength: EmptyEvalLength{
					Value: 37,
				},
				ErrorFF: ErrorFF{
					Value: false,
				},
				Chrome: Chrome{
					Value: false,
				},
				CookiesEnabled: CookiesEnabled{
					Value: true,
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewExperimentalInMemory()

			e.fingerprints = tt.fingerprints

			got, got1 := e.Do(tt.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Do() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExperimentalInMemory_Store(t *testing.T) {
	c := NewExperimentalInMemory()

	keyA, keyB, keyC := "a", "b", "c"

	if e, w := c.Store(keyA, Fingerprint{VisitorID: keyA}), 0.; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}

	if len(c.fingerprints) != 1 {
		t.Error("storage inconsistency")
	}
	if len(c.fingerprints[keyA]) != 1 {
		t.Errorf("storage inconsistency under key = %s", keyA)
	}
	if len(c.fingerprints[keyB]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyB)
	}
	if len(c.fingerprints[keyC]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyC)
	}

	if e, w := c.Store(keyA, Fingerprint{VisitorID: keyB}), 0.; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}

	if len(c.fingerprints) != 2 {
		t.Error("storage inconsistency")
	}
	if len(c.fingerprints[keyA]) != 2 {
		t.Errorf("storage inconsistency under key = %s", keyA)
	}
	if len(c.fingerprints[keyB]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyB)
	}
	if len(c.fingerprints[keyC]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyC)
	}

	x := -2./3*math.Log2(2./3) - 1./3*math.Log2(1./3)

	if e, w := c.Store(keyB, Fingerprint{VisitorID: keyC}), x; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}

	if len(c.fingerprints) != 3 {
		t.Error("storage inconsistency")
	}
	if len(c.fingerprints[keyA]) != 2 {
		t.Errorf("storage inconsistency under key = %s", keyA)
	}
	if len(c.fingerprints[keyB]) != 1 {
		t.Errorf("storage inconsistency under key = %s", keyB)
	}
	if len(c.fingerprints[keyC]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyC)
	}

	if e, w := c.Store(keyA, Fingerprint{VisitorID: keyA}), x; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}
	if e, w := c.Store(keyA, Fingerprint{VisitorID: keyB}), x; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}
	if e, w := c.Store(keyB, Fingerprint{VisitorID: keyC}), x; e != w {
		t.Errorf("wrong entropy value; got %f, want %f", e, w)
	}

	if len(c.fingerprints) != 3 {
		t.Error("storage inconsistency")
	}
	if len(c.fingerprints[keyA]) != 2 {
		t.Errorf("storage inconsistency under key = %s", keyA)
	}
	if len(c.fingerprints[keyB]) != 1 {
		t.Errorf("storage inconsistency under key = %s", keyB)
	}
	if len(c.fingerprints[keyC]) != 0 {
		t.Errorf("storage inconsistency under key = %s", keyC)
	}
}

func TestNewExperimentalInMemory(t *testing.T) {
	tests := []struct {
		name string
		want ExperimentalInMemory
	}{
		{
			name: "basic",
			want: ExperimentalInMemory{
				fingerprints: make(map[string][]Fingerprint),
				mtx:          &sync.RWMutex{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExperimentalInMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExperimentalInMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
