. {
    param (
        [Parameter()]
        [string]$turnsString
    )

    [int[]]$turns = $turnsString -split ", "
    [int[]]$firstPlayerTurns = @()
    [int[]]$secondPlayerTurns = @()

    $WinCellCombinations = @(
        @(0, 1, 2),
        @(3, 4, 5),
        @(6, 7, 8),
        @(0, 3, 6),
        @(1, 4, 7),
        @(2, 5, 8),
        @(0, 4, 8),
        @(2, 4, 6)    
    )
    
    for ($i = 0; $i -lt $turns.Count; $i++) {
        switch ($i % 2) {
            0 {
                $firstPlayerTurns += , $turns[$i]
                $WinCellCombinations.ForEach{
                    if ((([System.Linq.Enumerable]::Intersect($firstPlayerTurns, [int[]]$_)).foreach{ $_ }).Count -eq 3) {
                        Write-Host ("First player win at [{0}] turn" -f ($i + 1))
                        exit
                    }
                }
            }
            1 {
                $secondPlayerTurns += , $turns[$i] 
                $WinCellCombinations.ForEach{
                    if ((([System.Linq.Enumerable]::Intersect($secondPlayerTurns, [int[]]$_)).foreach{ $_ }).Count -eq 3) {
                        Write-Host ("Second player win at [{0}] turn" -f ($i + 1))
                        exit
                    }
                }
            }
        }
    }
    Write-Host "Draw"
 } -turnsString "7, 0, 5, 4, 2, 3, 1, 6, 8"
