#!/run/current-system/sw/bin/fish

echo "Testing remcom tool..."
echo "======================"

cd testdata

set -g tests_passed 0
set -g tests_failed 0
set -g total_tests 0

function test_file
    set input_file $argv[1]
    set test_name $argv[2]
    
    echo "Testing: $test_name"
    echo "Input: $input_file"
    
    cp $input_file $input_file.test
    
    ../remcom $input_file.test --mode=auto
    
    echo "✅ PROCESSED: $input_file"
    echo "Lines before and after processing:"
    wc -l $input_file $input_file.test | head -2
    
    echo "----------------------------------------"
    echo
end

function test_file_with_expected
    set input_file $argv[1]
    set expected_file $argv[2]
    set test_name $argv[3]
    
    set -g total_tests (math $total_tests + 1)
    
    echo "Testing: $test_name"
    echo "Input: $input_file"
    echo "Expected: $expected_file"
    
    cp $input_file $input_file.test
    
    ../remcom $input_file.test --mode=auto
    
    if test -f $expected_file
        if diff -u $expected_file $input_file.test > /dev/null
            echo "✅ PASS: Output matches expected result"
            set -g tests_passed (math $tests_passed + 1)
        else
            echo "❌ FAIL: Output differs from expected result"
            set -g tests_failed (math $tests_failed + 1)
            echo "Differences:"
            diff -u $expected_file $input_file.test
        end
    else
        echo "⚠️  No expected file found, showing processed result"
        echo "Processed file content:"
        head -10 $input_file.test
    end
    
    echo "----------------------------------------"
    echo
end

echo "Backing up original test files..."
set test_files basic_comments.go string_literals.go raw_strings.go escaped_characters.go mixed_content.go unicode_content.go nested_quotes.go edge_cases.go multiline_raw_string.go only_comments.txt single_line.go line_endings_crlf.txt empty_file.txt

for file in $test_files
    if test -f $file
        cp $file $file.orig
    end
end

echo "Running comprehensive test suite..."
echo

test_file_with_expected "basic_comments.go" "basic_comments.go.expected" "Basic Comments Removal"
test_file_with_expected "mixed_content.go" "mixed_content.go.expected" "Mixed Content with Blank Lines"
test_file_with_expected "single_line.go" "single_line.go.expected" "Single Line File"

test_file_with_expected "string_literals.go" "string_literals.go.expected" "Comments in String Literals (Should NOT Remove)"
test_file_with_expected "raw_strings.go" "raw_strings.go.expected" "Comments in Raw Strings (Should NOT Remove)"
test_file_with_expected "multiline_raw_string.go" "multiline_raw_string.go.expected" "Multiline Raw Strings with Fake Comments"

test_file_with_expected "escaped_characters.go" "escaped_characters.go.expected" "Escaped Characters and Quotes"
test_file_with_expected "nested_quotes.go" "nested_quotes.go.expected" "Complex Nested Quote Scenarios"
test_file_with_expected "edge_cases.go" "edge_cases.go.expected" "Various Edge Cases and Boundaries"
test_file_with_expected "unicode_content.go" "unicode_content.go.expected" "Unicode Characters and International Text"

test_file_with_expected "only_comments.txt" "only_comments.txt.expected" "File with Only Comments"
test_file_with_expected "empty_file.txt" "empty_file.txt.expected" "Empty File"
test_file_with_expected "line_endings_crlf.txt" "line_endings_crlf.txt.expected" "CRLF Line Endings"

echo "Testing complete!"
echo

echo "==============================="
echo "TEST SUMMARY"
echo "==============================="
echo "Total Tests Run: $total_tests"
echo "✅ Passed: $tests_passed"
echo "❌ Failed: $tests_failed"

if test $tests_failed -eq 0
    echo "🎉 ALL TESTS PASSED!"
    set pass_rate 100
else
    set pass_rate (math "$tests_passed * 100 / $total_tests")
    set pass_rate_decimal (printf "%.1f" (math "$tests_passed * 100 / $total_tests"))
    echo "📊 Pass Rate: $pass_rate_decimal%"
end
echo "==============================="
echo

echo "Restoring original test files..."
for file in $test_files
    if test -f $file.orig
        mv $file.orig $file
    end
end

echo "Cleaning up test copies..."
for file in $test_files
    if test -f $file.test
        rm $file.test
    end
end

echo "✅ Original test files restored and cleaned up"