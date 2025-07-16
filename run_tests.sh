#!/run/current-system/sw/bin/bash

echo "Testing remcom tool..."
echo "======================"

cd testdata

tests_passed=0
tests_failed=0
total_tests=0

function test_file() {
    local input_file="$1"
    local test_name="$2"
    
    echo "Testing: $test_name"
    echo "Input: $input_file"
    
    cp "$input_file" "$input_file.test"
    
    ../remcom "$input_file.test" --mode=auto
    
    echo "✅ PROCESSED: $input_file"
    echo "Lines before and after processing:"
    wc -l "$input_file" "$input_file.test" | head -2
    
    echo "----------------------------------------"
    echo
}

function test_file_with_expected() {
    local input_file="$1"
    local expected_file="$2"
    local test_name="$3"
    
    ((total_tests++))
    
    echo "Testing: $test_name"
    echo "Input: $input_file"
    echo "Expected: $expected_file"
    
    cp "$input_file" "$input_file.test"
    
    ../remcom "$input_file.test" --mode=auto
    
    if [[ -f "$expected_file" ]]; then
        if diff -u "$expected_file" "$input_file.test" > /dev/null; then
            echo "✅ PASS: Output matches expected result"
            ((tests_passed++))
        else
            echo "❌ FAIL: Output differs from expected result"
            ((tests_failed++))
            echo "Differences:"
            diff -u "$expected_file" "$input_file.test"
        fi
    else
        echo "⚠️  No expected file found, showing processed result"
        echo "Processed file content:"
        head -10 "$input_file.test"
    fi
    
    echo "----------------------------------------"
    echo
}

echo "Backing up original test files..."
test_files=(basic_comments.go string_literals.go raw_strings.go escaped_characters.go mixed_content.go unicode_content.go nested_quotes.go edge_cases.go multiline_raw_string.go only_comments.txt single_line.go line_endings_crlf.txt empty_file.txt)

for file in "${test_files[@]}"; do
    if [[ -f "$file" ]]; then
        cp "$file" "$file.orig"
    fi
done

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

if [[ $tests_failed -eq 0 ]]; then
    echo "🎉 ALL TESTS PASSED!"
    pass_rate=100
else
    pass_rate=$(awk "BEGIN {printf \"%.1f\", $tests_passed * 100 / $total_tests}")
    echo "📊 Pass Rate: $pass_rate%"
fi
echo "==============================="
echo

echo "Restoring original test files..."
for file in "${test_files[@]}"; do
    if [[ -f "$file.orig" ]]; then
        mv "$file.orig" "$file"
    fi
done

echo "Cleaning up test copies..."
for file in "${test_files[@]}"; do
    if [[ -f "$file.test" ]]; then
        rm "$file.test"
    fi
done

echo "✅ Original test files restored and cleaned up" 