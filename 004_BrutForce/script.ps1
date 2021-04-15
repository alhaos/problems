. {
    param (
        [Parameter(Position=0)]
        [string[]]$Elements,

        [Parameter(Position=1)]
        [int]$Positions
    )
    
    $positions -= 2
    $StackA = [System.Collections.Stack]::new()
    $StackB = [System.Collections.Stack]::new()
    
    $Elements[-1..0].ForEach{
        $StackA.Push(@($_))
    }

    (0..$positions).ForEach{
        while ($StackA.Count) {
            $CurrentCase = $StackA.Pop()
            $Elements.ForEach{
                $StackB.Push(($CurrentCase + , $_))
            }
        }

        while ($StackB.Count) {
            $StackA.Push($StackB.Pop())
        }
    }
    
    $StackA.ForEach{$_ -join ", "}
} -Elements @("[red]", "[grn]", "[blu]") -Positions 3


